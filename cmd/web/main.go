package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/dupe/view", dupeView)
	mux.HandleFunc("/dupe/create", dupeCreate)

	srv := &http.Server{
		Addr:     ":4000",
		ErrorLog: errorLog,
		Handler:  mux,
	}

	// Init web server
	infoLog.Print("Starting server on :4000")
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
