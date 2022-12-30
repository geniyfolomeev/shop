package main

import (
	"log"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
}

func main() {
	http.HandleFunc("/api/games", handler)

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
