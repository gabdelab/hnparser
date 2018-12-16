package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

func main() {
	port := flag.Int("port", 8080, "listening port")
	filePathPtr := flag.String("file-path", os.Getenv("HN_LOGS"), "path to HN logs")
	flag.Parse()

	logs := logs{}
	err := parseLogs(*filePathPtr, logs)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to parse logs file"))
	}

	handler := GetHandler(logs)

	http.HandleFunc("/1/count/", handler.Count)
	http.HandleFunc("/1/popular/", handler.Popular)
	fmt.Printf("server starting on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
