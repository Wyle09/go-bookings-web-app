package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/wyle09/go-bookings-web-app/pkg/config"
	"github.com/wyle09/go-bookings-web-app/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// Sets the config for the template package. Takes pointer of the app config
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(templateData *models.TemplateData) *models.TemplateData {
	return templateData
}

// Renders templates from the template folder
func RenderTemplate(writer http.ResponseWriter, tmpl string, templateData *models.TemplateData) {

	var templateCache map[string]*template.Template
	if app.UseCache {
		// Get the template cache from the app config
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	template, validTemplate := templateCache[tmpl]
	if !validTemplate {
		log.Fatal("Could not get template from template cache")
	}

	buffer := new(bytes.Buffer)
	templateData = AddDefaultData(templateData)
	_ = template.Execute(buffer, templateData)

	_, err := buffer.WriteTo(writer)
	if err != nil {
		log.Fatal(err)
	}
}

// Creates template cache as a map type.
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	// Get the page templates
	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = templateSet
	}

	return myCache, nil
}
