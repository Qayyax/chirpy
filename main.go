package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "8080"

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    "localhost:" + port,
		Handler: mux,
	}

	fileServer := http.FileServer(http.Dir("."))
	mux.Handle("/", fileServer)
	log.Printf("Serving on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
