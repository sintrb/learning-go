package main

import (
	"net/http"
	"io/ioutil"
)


func main() {
	resp, err := http.Get("http://www.baidu.com/")
	if err != nil{
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err  != nil{
		return
	}
	println(body)
}