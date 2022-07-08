package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, htmlDocument string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + htmlDocument)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
	}

}
