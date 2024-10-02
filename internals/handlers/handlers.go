package handlers

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"
	"sip/config"
	"sip/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	names := make(map[string]string)
	names["owner"] = "alex"
	renderTemplate(w, "home", &models.TemplateData{
		StringData: names,
	})
}

func Help(w http.ResponseWriter, r *http.Request) {
	numbers := make(map[string]int)
	numbers["Number"] = 10
	renderTemplate(w, "help", &models.TemplateData{
		IntData: numbers,
	})
}

var appConfig *config.Config

func CreateTemple(app *config.Config) {
	appConfig = app
}

func renderTemplate(w http.ResponseWriter, tmplName string, td *models.TemplateData) {
	templateCache := appConfig.TemplateCache
	tmpl, ok := templateCache[tmplName+".page.tmpl"]

	if !ok {
		http.Error(w, "le template n'éxiste pas", http.StatusInternalServerError)
		return
	}

	buffer := new(bytes.Buffer)
	tmpl.Execute(buffer, td)
	buffer.WriteTo(w)
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		tmpl := template.Must(template.ParseFiles(page))
		layouts, err := filepath.Glob("./templates/layouts/*.layout.tmpl")
		if err != nil {
			return cache, err
		}
		if len(layouts) > 0 {
			tmpl.ParseGlob("./templates/layouts/*.layout.tmpl")
		}
		cache[name] = tmpl
	}
	return cache, nil
}
