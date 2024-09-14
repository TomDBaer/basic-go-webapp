package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/TomDBaer/basic-go-webapp/pkg/config"
	"github.com/TomDBaer/basic-go-webapp/pkg/models"
)

// Simpel und funktioniert wenn man wenig templates hat und man keinen cache nutzen will
// RenderTemplateBasic renders the html templates || NOT USED
func RenderTemplateBasic(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("error parsing template:", err)
		return
	}
}

// Template render, cache und map4
var tc = make(map[string]*template.Template)

// RenderTemplate renders the html templates
func RenderTemplate(w http.ResponseWriter, templ string) {
	// adresse in tmpl schreiben
	var tmpl *template.Template
	var err error

	// check if the template is in the cache
	_, inMap := tc[templ]
	if !inMap {
		// not in cache, need to be created
		err = CreateTemplateCache(templ)
		if err != nil {
			log.Println(err)
		}
	} else {
		// template in cache
		log.Println("using cached template")
	}

	tmpl = tc[templ]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.html",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}

	// add template to cache (map)
	// bsp: map[home.page.html:0xc00008c330]
	// tmpl ist in der adresse
	tc[t] = tmpl

	return nil
}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// TODO For later use
func AddDefaultData(templateData *models.TemplateData) *models.TemplateData {

	return templateData
}

// RenderTemplateAdvanced renders the html templates || partialy used
func RenderTemplateAdvanced(w http.ResponseWriter, tmpl string, templateData *models.TemplateData) {

	// create a template chache
	// Hier wird der cache von der config cache geladen
	templateCache := app.TemplateCacheAdvanced

	// get request template from cache
	t, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("Could not get temlate from template cache")
	}

	// buffer || nicht nötig, soll bei der Fehlersuche in der map helfen
	buf := new(bytes.Buffer)

	templateData = AddDefaultData(templateData)

	// Wenn ich Daten übergeben will, ansonsten nil
	err := t.Execute(buf, templateData)
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
	// Wenn ich den buffer nicht nutzen will
	// err = t.Execute(w, nil)
	// if err != nil {
	// 	log.Println(err)
	// }
}

func CreateTemplateCacheAdvanced() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all files named *.page.html from ./templates
	pages, err := filepath.Glob("./templates/*.page.html")
	log.Println("pages:", pages)
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.html
	for _, page := range pages {
		name := filepath.Base(page)
		log.Println("page:", page)
		log.Println("name:", name)

		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = templateSet
	}
	return myCache, nil
}
