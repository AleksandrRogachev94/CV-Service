package main

import (
	"context"
	"encoding/json"
	fmt "fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func healthHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
}

func recognizeIndex() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get url params
		vars := mux.Vars(r)

		// respond with saved recognition.
		files := []FileLocation{
			FileLocation{Key: "key" + vars["id"], Bucket: "Bucket"},
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode([][]FileLocation{files}); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func recognizeShow() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get url params
		vars := mux.Vars(r)

		// respond with saved recognition.
		files := []FileLocation{
			FileLocation{Key: "key" + vars["id"], Bucket: "Bucket"},
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(files); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func recognize(grpcClient CVServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
		defer cancel()
		t, _ := ctx.Deadline()
		fmt.Println("--->", t)
		// get url params
		// vars := mux.Vars(r)

		// get todo from grpc channel
		file := FileLocation{Bucket: "go-cvservice", Key: "birds2.jpg"}
		files, err := grpcClient.Recognize(ctx, &RecognizeRequest{File: &file})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		fmt.Println(files)

		// respond with received todo
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(files); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
