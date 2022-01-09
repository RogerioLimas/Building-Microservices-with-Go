package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/RogerioLimas/Building-Microservices-with-Go/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api\t", log.LstdFlags)

	hh := handlers.NewHello(l)
	hg := handlers.NewGoodBye(l)

	
	mux := http.NewServeMux()
	mux.Handle("/", hh)
	mux.Handle("/goodbye", hg)
	
	srv := http.Server {
		Addr: ":8080",
		Handler: mux,
		ReadTimeout: 10 * time.Millisecond,
		WriteTimeout: 10 * time.Millisecond,
		IdleTimeout: 120 * time.Millisecond,
	}

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)


	sig := <- sigChan
	log.Println("Received terminate, graceful shutdown", sig)
	
	ctx, _ := context.WithTimeout(context.Background(), 40 * time.Second)
	srv.Shutdown(ctx)

}