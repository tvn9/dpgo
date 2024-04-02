package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type appData struct {
	Data map[string]any
}

func (a *application) render(w http.ResponseWriter, t string, td *appData) {
	var tmpl *template.Template

	// if we are using the template cache, try to get the template from map.
	if a.config.useCache {
		if tmplFromMap, ok := a.templateMap[t]; ok {
			tmpl = tmplFromMap
		}
	}

	if tmpl == nil {
		newTemplate, err := a.createTemplate(t)
		if err != nil {
			log.Println("Error building template:", err)
			return
		}
		fmt.Println("create template from disk")
		tmpl = newTemplate
	}
	if td == nil {
		td = &appData{}
	}

	if err := tmpl.ExecuteTemplate(w, t, td); err != nil {
		log.Fatalf("fails to execute template, %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a *application) createTemplate(t string) (*template.Template, error) {
	templateSlice := []string{
		"./templates/base.gohtml",
		"./templates/header.gohtml",
		"./templates/navbar.gohtml",
		"./templates/footer.gohtml",
		"./templates/index.gohtml",
		"./templates/contact.gohtml",
		"./templates/about.gohtml",
		"./templates/login.gohtml",
		"./templates/" + t,
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		return nil, err
	}

	a.templateMap[t] = tmpl
	return tmpl, nil
}
