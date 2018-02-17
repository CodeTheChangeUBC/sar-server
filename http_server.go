package main

import (
	"log"
	"net/http"
)

func runServer() {
	fs := http.FileServer(http.Dir("./map-display/dist"))
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
