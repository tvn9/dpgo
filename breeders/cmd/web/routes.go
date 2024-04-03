package main

import "net/http"

func (a *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./local/"))))

	mux.HandleFunc("/", a.Home)
	mux.HandleFunc("/about", a.About)
	mux.HandleFunc("/cat-breed", a.CatBreed)
	mux.HandleFunc("/cat-breeds", a.CatBreeds)
	mux.HandleFunc("/dog-breed", a.DogBreed)
	mux.HandleFunc("/dog-breeds", a.DogBreeds)

	return mux
}
