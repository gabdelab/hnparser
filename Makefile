build:
	dep ensure
	go build

run:
	./hnparser

test:
	go test ./...
