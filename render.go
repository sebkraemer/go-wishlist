package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemple, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemple.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
	}
}
