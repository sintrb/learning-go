package main

import (
	// "fmt"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("<h1>Hi it's HTTPS</h1>"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServeTLS(":8080", "./server.cert", "client.store", nil)
}
