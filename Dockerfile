FROM golang:latest

RUN mkdir -p /go/src/hnparser/
RUN go get -u github.com/golang/dep/cmd/dep

ADD . /go/src/hnparser/
RUN rm /go/src/hnparser/hnparser

WORKDIR /go/src/hnparser/

RUN dep ensure
RUN GOOS=linux go build -o hnparser .

CMD ["./hnparser", "-file-path", "/usr/share/results.tsv"]
