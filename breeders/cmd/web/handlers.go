package main

import (
	"net/http"
)

func (a *application) Home(w http.ResponseWriter, r *http.Request) {
	a.render(w, "index.html", nil)
}

func (a *application) Contact(w http.ResponseWriter, r *http.Request) {
	a.render(w, "contact.html", nil)
}

func (a *application) About(w http.ResponseWriter, r *http.Request) {
	a.render(w, "about.html", nil)
}

func (a *application) Login(w http.ResponseWriter, r *http.Request) {
	a.render(w, "login.html", nil)
}
