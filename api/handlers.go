package main

import (
	"context"
	"encoding/json"
	fmt "fmt"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type resultItem struct {
	Conf float64 `json:"conf"`
	URL  string  `json:"url"`
}

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

		ctx, cancel := context.WithTimeout(r.Context(), 20*time.Second)
		defer cancel()
		t, _ := ctx.Deadline()
		fmt.Println("--->", t)

		var data map[string]string
		err := json.NewDecoder(r.Body).Decode(&data)

		if data["bucket"] == "" || data["key"] == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// send request to from grpc channel
		file := FileLocation{Bucket: data["bucket"], Key: data["key"]}
		response, err := grpcClient.Recognize(ctx, &RecognizeRequest{File: &file})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		resultItemsDict := make(map[string][]resultItem)
		for _, item := range response.Items {
			resultItemsDict[item.Label] = append(resultItemsDict[item.Label], resultItem{
				URL:  "https://" + item.Location.Bucket + ".s3.amazonaws.com/" + item.Location.Key,
				Conf: item.Conf,
			})
		}

		// respond with received todo
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resultItemsDict); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func upload(uploader *s3manager.Uploader) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 10 mb
		r.Body = http.MaxBytesReader(w, r.Body, 10048576)
		bucket := "go-cvservice-assets"
		key := (uuid.New()).String() + ".jpg"
		acl := "public-read"
		contentType := "image/jpeg"
		_, err := uploader.Upload(&s3manager.UploadInput{
			Bucket:      aws.String(bucket),
			Key:         aws.String(key),
			Body:        r.Body,
			ACL:         &acl,
			ContentType: &contentType,
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		file := FileLocation{Key: key, Bucket: bucket}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(file); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
