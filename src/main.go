package main

import (
	"net/http"
	"log"
)

func main() {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
