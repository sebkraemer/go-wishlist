package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/sebkraemer/go-wishlist/pkg/config"
	"github.com/sebkraemer/go-wishlist/pkg/models"
)

// functions for later use in template parsing
var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplate sets the config for the render package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData adds data that shall be available to all pages
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("template not found in cache ", tmpl)
	}

	buf := new(bytes.Buffer) // use a buffer to catch execute error early (extra step)

	td = AddDefaultData(td)
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*page.tmpl")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		matches, err := filepath.Glob("./templates/*layout.tmpl")
		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return cache, err
			}
			cache[name] = ts
		}
	}

	return cache, nil
}
