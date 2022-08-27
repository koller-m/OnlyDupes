package main

import (
	"log"
	"net/http"
)

// Create home handler
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from OnlyDupes"))
}

func main() {
	// Init router, then register the home handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Init web server
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
