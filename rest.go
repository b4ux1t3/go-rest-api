package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/b4ux1t3/libgollatz"
)

var templates = template.Must(template.ParseFiles("result.html"))

func collatzHandler(w http.ResponseWriter, r *http.Request) {
	argumentIndex := len("/collatz/")
	argument, err := strconv.ParseUint(r.URL.Path[argumentIndex:], 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := libgollatz.Collatz(argument)

	renderTemplate(w, result)
}

func renderTemplate(w http.ResponseWriter, r libgollatz.Result) {
	err := templates.ExecuteTemplate(w, "result.html", r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/collatz/", collatzHandler)

	http.ListenAndServe(":8080", nil)
}
