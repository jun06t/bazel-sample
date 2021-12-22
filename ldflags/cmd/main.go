package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jun06t/bazel-sample/ldflags/alive"
)

func main() {
	fmt.Println("v5")
	http.HandleFunc("/alive", alive.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
