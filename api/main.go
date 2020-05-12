package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	grpc "google.golang.org/grpc"
)

func main() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	logger.Println("Connecting to grpc...")
	dialOpts := []grpc.DialOption{grpc.WithBlock(), grpc.WithInsecure()}
	conn, err := grpc.Dial(viper.Get("PROCESSOR_ADDRESS").(string), dialOpts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := NewCVServiceClient(conn)
	logger.Println("Connected to grpc")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(viper.Get("AWS_ACCESS_KEY").(string), viper.Get("AWS_SECRET_KEY").(string), ""),
	})
	if err != nil {
		logger.Fatalf("failed to configure aws: %v", err)
	}
	uploader := s3manager.NewUploader(sess)

	port := "8080"
	router := mux.NewRouter()

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      handlers.LoggingHandler(os.Stdout, router),
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 20 * time.Second,
		// WriteTimeout: 1 * time.Millisecond,
		IdleTimeout: 15 * time.Second,
	}

	router.Handle("/health", healthHandler()).Methods("GET")
	router.Handle("/recognitions", recognize(client)).Methods("POST")
	router.Handle("/recognitions", recognizeIndex()).Methods("GET")
	router.Handle("/recognitions/{id}", recognizeShow()).Methods("GET")
	router.Handle("/upload", upload(uploader)).Methods("POST")

	logger.Printf("Listening on port %s...", port)
	server.ListenAndServe()
}
