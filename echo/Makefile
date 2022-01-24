.PHONY: modsync
modsync:
	go get -u
	go mod tidy
	go mod vendor

.PHONY: docker
docker: modsync
	docker build -t echo-demoapp . --no-cache
	docker run --rm -it -p 8080:8080 echo-demoapp:latest
