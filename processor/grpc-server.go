package main

import (
	context "context"
	fmt "fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/spf13/viper"
)

type cvServiceServer struct {
	UnimplementedCVServiceServer
	rekognition *rekognition.Rekognition
}

func newServer() *cvServiceServer {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-2"),
		Credentials: credentials.NewStaticCredentials(viper.Get("AWS_ACCESS_KEY").(string), viper.Get("AWS_SECRET").(string), ""),
	})
	if err != nil {
		log.Fatalf("failed to configure aws: %v", err)
	}
	svc := rekognition.New(sess)

	return &cvServiceServer{rekognition: svc}
}

func (s *cvServiceServer) Recognize(c context.Context, r *RecognizeRequest) (*RecognizeResponse, error) {
	fmt.Println("GetTodo request")
	files := []*FileLocation{
		&FileLocation{Bucket: "bucket", Key: "key"},
	}
	resPt := &RecognizeResponse{Files: files}

	fmt.Println(r.File.Bucket, r.File.Key)

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

	return resPt, nil
}
