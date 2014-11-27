package main

import (
	"net/http"
	"fmt"
	"log"
)


func myHandler(w http.ResponseWriter, r *http.Request) {
	
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
