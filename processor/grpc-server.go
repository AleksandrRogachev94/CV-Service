package main

import (
	"bytes"
	context "context"
	fmt "fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/viper"
)

// Helper function to measure execution time
func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

// Result of processing single recognize response instance
type processInstanceResult struct {
	bucket string
	key    string
	err    error
}

// Struct holding grcp server instance. Includes instantiated aws-related services.
type cvServiceServer struct {
	UnimplementedCVServiceServer
	rekognition *rekognition.Rekognition
	downloader  *s3manager.Downloader
	uploader    *s3manager.Uploader
}

// Create new grpc server instance
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

// Recognize labels in the provided image and extract + upload bounding boxes.
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

// Download source file from S3 and convert it to in-memory jpeg image.
func (s *cvServiceServer) downloadSource(bucket string, key string) (*image.Image, error) {
	// use temporary buffer to store download results
	var awsBuff aws.WriteAtBuffer
	_, err := s.downloader.Download(&awsBuff,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		})
	// use temporary buffer to stream donwload results to image instance.
	srcBuff := bytes.NewReader(awsBuff.Bytes())
	src, err := jpeg.Decode(srcBuff)

	if err != nil {
		return nil, err
	}
	return &src, nil
}

// Get subImage by cropping src within box
func (s *cvServiceServer) getSubImage(src *image.Image, box rekognition.BoundingBox, dx float64, dy float64) (*image.Image, error) {
	rgbSrc := (*src).(interface {
		SubImage(r image.Rectangle) image.Image
	})
	x0 := int(*box.Left * dx)
	y0 := int(*box.Top * dy)
	xn := int((*box.Left + *box.Width) * dx)
	yn := int((*box.Top + *box.Height) * dy)
	subImage := rgbSrc.SubImage(image.Rect(x0, y0, xn, yn))
	return &subImage, nil
}

// Upload given image to s3 bucket
func (s *cvServiceServer) uploadImage(image *image.Image, bucket string, key string) error {
	var subBuff bytes.Buffer
	png.Encode(&subBuff, *image)
	acl := "public-read"
	contentType := "image/png"
	_, err := s.uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        &subBuff,
		ACL:         &acl,
		ContentType: &contentType,
	})
	return err
}

// Use recognize result instance to extract subImage from src and upload it to s3
func (s *cvServiceServer) processInstance(
	src *image.Image,
	inst *rekognition.Instance,
	label *rekognition.Label,
	prefix int,
	bucket string,
	key string,
	ch chan processInstanceResult,
) {
	dx := float64((*src).Bounds().Dx())
	dy := float64((*src).Bounds().Dy())
	subImage, err := s.getSubImage(src, *inst.BoundingBox, dx, dy)
	if err != nil {
		ch <- processInstanceResult{bucket: "", key: "", err: err}
		return
	}

	uploadKey := fmt.Sprintf("%s-results/%s-%.2f/%d-%.2f.png", key, *label.Name, *label.Confidence, prefix, *inst.Confidence)
	err = s.uploadImage(subImage, bucket, uploadKey)
	if err != nil {
		ch <- processInstanceResult{bucket: "", key: "", err: err}
		return
	}

	ch <- processInstanceResult{bucket: bucket, key: uploadKey, err: nil}
}

// Process results returned from aws rekognize service
func (s *cvServiceServer) processRecognizeResult(c context.Context, bucket string, key string, r *rekognition.DetectLabelsOutput) ([]*FileLocation, error) {
	defer elapsed("processRecognizeResult")()
	var files []*FileLocation

	src, err := s.downloadSource(bucket, key)
	if err != nil {
		return nil, err
	}

	ch := make(chan processInstanceResult)

	cnt := 0
	for _, label := range r.Labels {
		for j, inst := range label.Instances {
			go s.processInstance(src, inst, label, j, bucket, key, ch)
			cnt++
		}
	}
	for i := 0; i < cnt; i++ {
		instRes := <-ch
		if instRes.err != nil {
			return nil, err
		}
		files = append(files, &FileLocation{Bucket: instRes.bucket, Key: instRes.key})
	}
	return files, nil
}
