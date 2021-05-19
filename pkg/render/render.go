package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/kjuchniewicz/go-web-app/models"
	"github.com/kjuchniewicz/go-web-app/pkg/config"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates ustawia konfigurację dla pakietu szablonów
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renderuje szablon za pomocą html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Nie można pobrać szablonu z pamięci podręcznej szablonów")
	}

	buff := new(bytes.Buffer)

	td = AddDefaultData(td)
	_ = t.Execute(buff, td)

	_, err := buff.WriteTo(w)
	if err != nil {
		fmt.Println("Błąd podczas zapisywania szablonu do przeglądarki")
	}
}

// CreateTemplateCache tworzy cache szablonów jako map'e
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}

		}

		myCache[name] = ts
	}
	return myCache, nil
}
