FROM golang:1.17-alpine3.14 AS builder
ENV GOOS=linux
ENV GOARCH=amd64
COPY . /data
WORKDIR /data
RUN go mod tidy && go build -o app main.go

FROM alpine:3.14
COPY --from=builder /data/app ./
EXPOSE 8080
CMD ["/app"]