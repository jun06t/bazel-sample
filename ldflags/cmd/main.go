package main

import (
	"log"
	"net/http"

	"github.com/jun06t/bazel-sample/ldflags/alive"
)

func main() {
	http.HandleFunc("/alive", alive.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
