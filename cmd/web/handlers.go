package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.html.render(w, http.StatusOK, nil, "base", "pages/home.tmpl")
}

func (app *application) chat(w http.ResponseWriter, r *http.Request) {
	app.html.render(w, http.StatusOK, nil, "base", "pages/chat.tmpl")
}
