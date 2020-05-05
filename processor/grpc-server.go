package main

import (
	"bytes"
	context "context"
	fmt "fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/viper"
)

type cvServiceServer struct {
	UnimplementedCVServiceServer
	rekognition *rekognition.Rekognition
	downloader  *s3manager.Downloader
	uploader    *s3manager.Uploader
}

func newServer() *cvServiceServer {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(viper.Get("AWS_ACCESS_KEY").(string), viper.Get("AWS_SECRET").(string), ""),
	})
	if err != nil {
		log.Fatalf("failed to configure aws: %v", err)
	}
	svc := rekognition.New(sess)
	downloader := s3manager.NewDownloader(sess)
	uploader := s3manager.NewUploader(sess)

	return &cvServiceServer{rekognition: svc, downloader: downloader, uploader: uploader}
}

func (s *cvServiceServer) Recognize(ctx context.Context, r *RecognizeRequest) (*RecognizeResponse, error) {
	input := &rekognition.DetectLabelsInput{
		Image: &rekognition.Image{
			S3Object: &rekognition.S3Object{
				Bucket: aws.String(r.File.Bucket),
				Name:   aws.String(r.File.Key),
			},
		},
		MaxLabels:     aws.Int64(123),
		MinConfidence: aws.Float64(70.000000),
	}
	result, err := s.rekognition.DetectLabels(input)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(result)

	files, err := s.processRecognizeResult(ctx, r.File.Bucket, r.File.Key, result)
	if err != nil {
		return nil, err
	}
	resPt := &RecognizeResponse{Files: files}
	return resPt, nil
}

func (s *cvServiceServer) processRecognizeResult(c context.Context, bucket string, key string, r *rekognition.DetectLabelsOutput) ([]*FileLocation, error) {
	var files []*FileLocation

	var awsBuff aws.WriteAtBuffer
	_, err := s.downloader.Download(&awsBuff,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		})
	if err != nil {
		return nil, err
	}

	srcBuff := bytes.NewReader(awsBuff.Bytes())
	src, err := jpeg.Decode(srcBuff)
	if err != nil {
		return nil, err
	}

	dx := float64(src.Bounds().Dx())
	dy := float64(src.Bounds().Dy())

	for _, label := range r.Labels {
		for j, inst := range label.Instances {
			rgbSrc := src.(interface {
				SubImage(r image.Rectangle) image.Image
			})
			x0 := int(*inst.BoundingBox.Left * dx)
			y0 := int(*inst.BoundingBox.Top * dy)
			xn := int((*inst.BoundingBox.Left + *inst.BoundingBox.Width) * dx)
			yn := int((*inst.BoundingBox.Top + *inst.BoundingBox.Height) * dy)
			fmt.Println(x0, y0, xn, yn)
			subImage := rgbSrc.SubImage(image.Rect(x0, y0, xn, yn))

			var subBuff bytes.Buffer
			png.Encode(&subBuff, subImage)
			acl := "public-read"
			contentType := "image/png"
			uploadKey := fmt.Sprintf("%s-results/%s-%.2f/%d-%.2f.png", key, *label.Name, *label.Confidence, j, *inst.Confidence)
			_, err := s.uploader.Upload(&s3manager.UploadInput{
				Bucket:      aws.String(bucket),
				Key:         aws.String(uploadKey),
				Body:        &subBuff,
				ACL:         &acl,
				ContentType: &contentType,
			})
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			files = append(files, &FileLocation{Bucket: bucket, Key: uploadKey})
		}
	}
	return files, nil
}
