package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fis, err := ioutil.ReadDir("./")
	fls := []string{}
	for _, fi := range fis {
		fls = append(fls, fi.Name())
		println(fi.Name())
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t, err := template.ParseFiles("tmp.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := make(map[string]interface{})
	data["files"] = fis
	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
