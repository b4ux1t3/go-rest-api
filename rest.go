package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/b4ux1t3/libgollatz"
)

// RequestNumber doesn't need to exist, I'm just having issues figuring something out
//TODO: Get rid of this
type RequestNumber struct {
	Num int
}

var templates = template.Must(template.ParseFiles("result.html", "index.html"))

func collatzHandler(w http.ResponseWriter, r *http.Request) {
	argumentIndex := len("/collatz/")
	argument, err := strconv.ParseUint(r.URL.Path[argumentIndex:], 10, 64)

	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		//http.Error(w, err.Error(), http.StatusInternalServerError)
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

func rootHandler(w http.ResponseWriter, r *http.Request) {
	var num RequestNumber
	num.Num = 1337
	err := templates.ExecuteTemplate(w, "index.html", num.Num)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func main() {
	http.HandleFunc("/collatz/", collatzHandler)
	http.HandleFunc("/", rootHandler)

	http.ListenAndServe(":8080", nil)
}
