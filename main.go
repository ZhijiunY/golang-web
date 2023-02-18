package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a Home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, fmt.Sprintf("This is a About page, and 2+2 = %d", AddValue(2, 2)))
}

func AddValue(x, y int) int {
	sum := x + y
	return sum
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValue(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "cannot divide by ZERO\n")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divide by %f is %f", 100.0, 0.0, f))
}

func divideValue(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide zero\n")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)
	http.ListenAndServe(portNumber, nil)
}
