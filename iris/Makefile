.PHONY: run
run:
	go run .

.PHONY: modsync
modsync:
	go get -u
	go mod tidy
	go mod vendor

.PHONY: docker
docker: modsync
	docker build -t iris-demoapp .
	docker run --rm -it -p 8080:8080 iris-demoapp:latest
