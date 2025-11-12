package main

import (
	"embed"
	"html/template"
	"net/http"
	"path/filepath"
)

//go:embed static/*
var staticFolder embed.FS

//go:embed templates/*
var templateFolder embed.FS

type pageData struct {
	PageTitle string
}

func LayoutedTemplate(path string) (tmpl *template.Template) {
	tmpl = template.Must(template.ParseFS(templateFolder, filepath.Join("templates", "layout.html"), path))
	return tmpl
}

func indexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := LayoutedTemplate(filepath.Join("templates", "main.html"))
		data := pageData{
			PageTitle: "Ganz viele Rezepte",
		}
		tmpl.Execute(w, data)
	}
}

func main() {
	router := http.NewServeMux()
	router.Handle("/", indexHandler())

	staticServer := http.FileServer(http.FS(staticFolder))
	router.Handle("/static/", staticServer)
	http.ListenAndServe(":3000", router)
}
