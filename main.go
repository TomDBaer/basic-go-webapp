package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const portNumber = ":8080"

// Handler Funktionen

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")
}

// renderTemplate renders the html templates
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("error parsing template:", err)
		return
	}
}

// Divide Route is just a placeholder.
func Divide(w http.ResponseWriter, r *http.Request) {
	x := 100.2
	y := 10.3
	f, err := divideValues(x, y)

	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	fmt.Fprintf(w, "%.2f divided by %.2f is %.2f", x, y, f)
}

// addValues adds two integers and returns the sum
func addValues(x, y int) int {
	return x + y
}

// divideValues divides two floats and returns the sum or error
func divideValues(x, y float64) (float64, error) {

	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}

	result := x / y
	return result, nil
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf("Starting application on port %s", portNumber)
	http.ListenAndServe(portNumber, nil)

}
