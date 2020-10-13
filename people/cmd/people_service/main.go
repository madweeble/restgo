package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	personHandler := func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Hello world")
	}

	http.HandleFunc("/person", personHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
