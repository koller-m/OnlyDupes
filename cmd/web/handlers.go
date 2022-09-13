package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/koller-m/OnlyDupes/internal/models"

	"github.com/julienschmidt/httprouter"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	dupes, err := app.dupes.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Dupes = dupes

	app.render(w, http.StatusOK, "home.tmpl.html", data)
}

func (app *application) dupeView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	dupe, err := app.dupes.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Dupe = dupe

	app.render(w, http.StatusOK, "view.tmpl.html", data)
}

func (app *application) dupeCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display the form for creating a new dupe..."))
}

func (app *application) dupeCreatePost(w http.ResponseWriter, r *http.Request) {
	dupe := "A dupe you never heard of!"
	content := "Another placeholder dupe while I finish the application"
	expires := 7

	id, err := app.dupes.Insert(dupe, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/dupe/view/%d", id), http.StatusSeeOther)
}
