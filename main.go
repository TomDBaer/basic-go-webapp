package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World!")
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("Number of bytes written: %d\n", n)
	})

	http.ListenAndServe(":8080", nil)

}
