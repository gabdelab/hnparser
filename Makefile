build:
	dep ensure
	go build

docker-build:
	docker build -t hnparser .

run:
	./hnparser

docker-run:
	docker run -v ~/Downloads/hn_logs.tsv:/usr/share/results.tsv -p 8080:8080 hnparser

test:
	go test ./...
