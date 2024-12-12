package main

import (
	_ "embed"
	"log"
)

//go:embed files/hello.txt
var hello string

func main() {
	log.Println(hello)
}
