# docker build -t demoapp . 
# docker run --rm -it -p 8080:8080 demoapp:latest
FROM golang:latest AS builder
RUN apt-get update
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /go/src/demoapp
COPY go.mod .
RUN go mod download
COPY . .
RUN go install

FROM scratch
COPY --from=builder /go/bin/demoapp .
ENTRYPOINT ["./demoapp"]
