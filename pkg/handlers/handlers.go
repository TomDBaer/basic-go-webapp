package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/TomDBaer/basic-go-webapp/pkg/render"
)

// Handler Funktionen

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.html")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.html")
}

//TODO: Kalkulation muss noch in eine eigene Datei geschoben werden

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
