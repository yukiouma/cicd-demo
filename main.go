package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(greet()))
	})
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func greet() string {
	return `{"msg": "hello, world"}`
}
