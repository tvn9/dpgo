package main

import "net/http"

func (a *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./local/"))))
	mux.HandleFunc("/", a.Home)
	mux.HandleFunc("/contact", a.Contact)
	mux.HandleFunc("/about", a.About)
	mux.HandleFunc("/login", a.Login)

	return mux
}
