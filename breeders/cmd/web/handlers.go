package main

import (
	"net/http"
)

func (a *application) Home(w http.ResponseWriter, r *http.Request) {
	a.render(w, "index.gohtml", nil)
}

func (a *application) Contact(w http.ResponseWriter, r *http.Request) {
	a.render(w, "contact.gohtml", nil)
}

func (a *application) About(w http.ResponseWriter, r *http.Request) {
	a.render(w, "about.gohtml", nil)
}

func (a *application) Login(w http.ResponseWriter, r *http.Request) {
	a.render(w, "login.gohtml", nil)
}
