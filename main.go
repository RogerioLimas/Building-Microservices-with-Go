package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.Hello)

	http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Good bye!")
		io.WriteString(rw, "Good bye!")
	})

	http.ListenAndServe(":8080", nil)
}