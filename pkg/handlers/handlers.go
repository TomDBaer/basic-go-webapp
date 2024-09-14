package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/TomDBaer/basic-go-webapp/pkg/config"
	"github.com/TomDBaer/basic-go-webapp/pkg/models"
	"github.com/TomDBaer/basic-go-webapp/pkg/render"
)

// Handler Funktionen

// Repository pattern wird hier verwendet
// Noch bescheiben was dies macht
// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// render.RenderTemplateAdvanced(w, "home.page.html")
	render.RenderTemplate(w, "home.page.html")
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello over there."

	// render.RenderTemplate(w, "about.page.html")

	render.RenderTemplateAdvanced(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
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
// func addValues(x, y int) int {
// 	return x + y
// }

// divideValues divides two floats and returns the sum or error
func divideValues(x, y float64) (float64, error) {

	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}

	result := x / y
	return result, nil
}
