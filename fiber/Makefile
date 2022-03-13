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
