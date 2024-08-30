package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Simpel und funktioniert wenn man wenig templates hat und man keinen cache nutzen will
// RenderTemplate renders the html templates
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
