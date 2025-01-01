package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the homae page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Home Page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the About page and 2 + 2 is %d", sum))
}

// addValues adds two ints and returns sum
func addValues(x, y int) int {
	return x + y
}

// main is the program entry point
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)

}
