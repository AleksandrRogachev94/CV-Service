FROM golang:alpine as builder
WORKDIR '/app'

# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM alpine:latest
COPY --from=builder /app/api /app/
CMD ["./app/api"]