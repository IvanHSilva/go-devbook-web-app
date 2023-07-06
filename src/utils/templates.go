package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

func ExecuteTemplate(w http.ResponseWriter, templeteFile string, data interface{}) {
	templates.ExecuteTemplate(w, templeteFile, data)
}
