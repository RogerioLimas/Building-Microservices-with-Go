package main

import (
	"log"
	"net/http"
	"os"

	"github.com/RogerioLimas/Building-Microservices-with-Go/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api\t", log.LstdFlags)

	hh := handlers.NewHello(l)
	hg := handlers.NewGoodBye(l)

	mux := http.NewServeMux()
	mux.Handle("/", hh)
	mux.Handle("/goodbye", hg)

	http.ListenAndServe(":8080", mux)
}