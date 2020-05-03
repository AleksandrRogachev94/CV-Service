package main

import (
	fmt "fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Connecting to grpc...")
	// dialOpts := []grpc.DialOption{grpc.WithBlock(), grpc.WithInsecure()}
	// conn, err := grpc.Dial(":8082", dialOpts...)
	// if err != nil {
	// 	log.Fatalf("fail to dial: %v", err)
	// }
	// defer conn.Close()
	fmt.Println("Connected to grpc")
	// client := NewCVServiceClient(conn)

	port := "8081"
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
	// router.Handle("/recognitions", recognize(client)).Methods("POST")
	router.Handle("/recognitions", recognizeIndex()).Methods("GET")
	router.Handle("/recognitions/{id}", recognizeShow()).Methods("GET")

	logger.Printf("Listening on port %s...", port)
	server.ListenAndServe()
}
