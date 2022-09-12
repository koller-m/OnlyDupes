package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/dupe/view", app.dupeView)
	mux.HandleFunc("/dupe/create", app.dupeCreate)

	return app.logRequest(secureHeaders(mux))
}
