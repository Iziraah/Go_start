package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

const addr string = "localhost:8080"

 func main() {

	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(addr, nil))

 }

 func handle(w http.ResponseWriter, r *http.Request) {
	
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	text := string(bodyBytes)
	responce := "1 instance" + text

	if _, err := w.Write([]byte(responce)); err != nil{
		log.Fatal(err)
	}

 }