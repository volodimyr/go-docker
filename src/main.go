package main

import (
	"log"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}

//The simplest one
//docker build -t go-docker-dev .
//docker run --rm -it -p 8080:8080 go-docker-dev

//Package management and layering
//docker build -t go-docker-dev .
//docker run --rm -it -v $(pwd):/go/src/app go-docker-dev bash

//Single Stage Production Build
//docker build -t go-docker-prod .
//docker run --rm -it -p 8080:8080 go-docker-prod
