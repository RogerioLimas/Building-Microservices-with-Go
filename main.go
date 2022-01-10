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

	// Create the handler
	hp := handlers.NewProducts(l)

	// Create a new server mux and register the handlers
	mux := http.NewServeMux()
	mux.Handle("/", hp)

	// Create a new server
	srv := http.Server{
		Addr:         ":8080",                // Bind the address
		Handler:      mux,                    // Set the default handler
		ErrorLog:     l,                      // Set the logger for the server
		ReadTimeout:  10 * time.Millisecond,  // Max time to read request from the client
		WriteTimeout: 10 * time.Millisecond,  // Max time to write response to the server
		IdleTimeout:  120 * time.Millisecond, // Set the idle timeout for connections using TCP Keep-Alive 
	}

	// Start the server
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	// Trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Received terminate, graceful shutdown", sig)

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)

}
