package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// Renders HTML template files for go to use.
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Errorf("error parsing template:", err)
		return
	}
}
