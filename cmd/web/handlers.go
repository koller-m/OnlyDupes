package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/koller-m/OnlyDupes/internal/models"
	"github.com/koller-m/OnlyDupes/internal/validator"

	"github.com/julienschmidt/httprouter"
)

type dupeCreateForm struct {
	Dupe                string `form:"dupe"`
	Content             string `form:"content"`
	Expires             int    `form:"expires"`
	validator.Validator `form:"-"`
}

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
	data := app.newTemplateData(r)

	data.Form = dupeCreateForm{
		Expires: 365,
	}

	app.render(w, http.StatusOK, "create.tmpl.html", data)
}

func (app *application) dupeCreatePost(w http.ResponseWriter, r *http.Request) {
	var form dupeCreateForm

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.Dupe), "dupe", "This field cannot be blank")
	form.CheckField(validator.MaxChars(form.Dupe, 100), "dupe", "This field cannot be more than 100 characters long")
	form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")
	form.CheckField(validator.PermittedInt(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7 or 365")

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "create.tmpl.html", data)
		return
	}

	id, err := app.dupes.Insert(form.Dupe, form.Content, form.Expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/dupe/view/%d", id), http.StatusSeeOther)
}
