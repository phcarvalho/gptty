package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello World</h1>")
}

func (app *application) systemsList(w http.ResponseWriter, r *http.Request) {
	tp, err := template.ParseFiles("./ui/html/pages/systemsList.tmpl")
	if err != nil {
		log.Print(err.Error())
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
	}

	systems, err := app.systems.List()
	if err != nil {
		log.Print(err.Error())
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
	}

	err = tp.Execute(w, systems)
	if err != nil {
		log.Print(err.Error())
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
	}
}
