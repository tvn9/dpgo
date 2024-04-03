package main

import (
	"net/http"
)

func (a *application) Home(w http.ResponseWriter, r *http.Request) {
	a.render(w, "index.html", nil)
}

func (a *application) About(w http.ResponseWriter, r *http.Request) {
	a.render(w, "about.html", nil)
}
func (a *application) CatBreed(w http.ResponseWriter, r *http.Request) {
	a.render(w, "cat-breed.html", nil)
}
func (a *application) CatBreeds(w http.ResponseWriter, r *http.Request) {
	a.render(w, "cat-breeds.html", nil)
}
func (a *application) DogBreed(w http.ResponseWriter, r *http.Request) {
	a.render(w, "dog-breed.html", nil)
}
func (a *application) DogBreeds(w http.ResponseWriter, r *http.Request) {
	a.render(w, "dog-breeds.html", nil)
}
