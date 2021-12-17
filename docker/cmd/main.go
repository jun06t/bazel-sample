package main

import (
	"log"

	"github.com/jun06t/bazel-sample/docker/uuid"
)

func main() {
	id, err := uuid.Generate()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(id)
}
