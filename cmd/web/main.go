package main

import (
	"fmt"
	"net/http"

	"github.com/TomDBaer/basic-go-webapp/pkg/handlers"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/divide", handlers.Divide)

	fmt.Printf("Starting application on port %s", portNumber)
	http.ListenAndServe(portNumber, nil)

}
