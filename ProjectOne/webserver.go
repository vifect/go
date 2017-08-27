package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadFile(r.URL.Path[1:])
	if err != nil {
		//body, _ = ioutil.ReadFile("index.html")
		log.Print(err)
		return
	}
	w.Write(body)
}
