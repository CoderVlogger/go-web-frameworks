.PHONY: run
run:
	go run cmd/demoapp/main.go

.PHONY: modsync
modsync:
	go get -u
	go mod tidy
	go mod vendor

.PHONY: docker
docker: modsync
	docker build -t echo-demoapp .
	docker run --rm -it -p 8080:8080 echo-demoapp:latest

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: format
format:
	find . -type f -name '*.go' -not -path './vendor/*' -exec gofmt -s -w {} +
	find . -type f -name '*.go' -not -path './vendor/*' -exec goimports -w {} +

.PHONY: test
test:
	go test ./... -race -cover
