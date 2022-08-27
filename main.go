package main

import (
	"log"
	"net/http"
)

// Create home handler
func home(w http.ResponseWriter, r *http.Request) {
	// Check if the request matches "/" exactly
	// If not, send a 404 response
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from OnlyDupes"))
}

func dupeView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific dupe..."))
}

func dupeCreate(w http.ResponseWriter, r *http.Request) {
	// Only allow dupeCreate to act on POST requests
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Add a new dupe..."))
}

func main() {
	// Init router, then register handler functions
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/dupe/view", dupeView)
	mux.HandleFunc("/dupe/create", dupeCreate)

	// Init web server
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
