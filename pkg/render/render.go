package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
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

// Template render, cache und map
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
		err = createTemplateCache(templ)
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

func createTemplateCache(t string) error {
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

// RenderTemplateAdvanced renders the html templates || NOT USED
func RenderTemplateAdvanced(w http.ResponseWriter, tmpl string) {

	// create a template chache
	templateCache, err := createTemplateCacheAdvanced()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("templateCache:", templateCache)
	// get request template from cache
	t, ok := templateCache[tmpl]
	if !ok {
		log.Fatal(err)
	}

	// buffer || nicht nötig, soll bei Fehlersuche in der map helfen
	buf := new(bytes.Buffer)

	err = t.Execute(buf, nil)
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

func createTemplateCacheAdvanced() (map[string]*template.Template, error) {
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
