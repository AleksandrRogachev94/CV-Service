package main

import (
	fmt "fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	grpc "google.golang.org/grpc"
)

func main() {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	fmt.Println("Connecting to grpc...")
	dialOpts := []grpc.DialOption{grpc.WithBlock(), grpc.WithInsecure()}
	conn, err := grpc.Dial(viper.Get("PROCESSOR_ADDRESS").(string), dialOpts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := NewCVServiceClient(conn)
	fmt.Println("Connected to grpc")

	port := viper.Get("PORT").(string)
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	router := mux.NewRouter()

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      handlers.LoggingHandler(os.Stdout, router),
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		// WriteTimeout: 1 * time.Millisecond,
		IdleTimeout: 15 * time.Second,
	}

	router.Handle("/health", healthHandler()).Methods("GET")
	router.Handle("/recognitions", recognize(client)).Methods("POST")
	router.Handle("/recognitions", recognizeIndex()).Methods("GET")
	router.Handle("/recognitions/{id}", recognizeShow()).Methods("GET")

	logger.Printf("Listening on port %s...", port)
	server.ListenAndServe()
}
