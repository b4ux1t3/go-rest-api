package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/b4ux1t3/libgollatz"
)

var templates = template.Must(template.ParseFiles("result.html", "index.html"))

func collatzHandler(w http.ResponseWriter, r *http.Request) {
	argumentIndex := len("/collatz/")
	argument, err := strconv.ParseUint(r.URL.Path[argumentIndex:], 10, 64)

	// This executes if we do not receive an integer from the user after "/collatz/" in the URL path.
	if err != nil {
		//http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		http.Error(w, "Unable to parse argument. URL path should be: /collatz/<POSITIVE BASE 10 INTEGER>", http.StatusInternalServerError)
		return
	}

	// No matter what, we want to run Collatz on the argument for now.
	// When we start doing result caching, though, we will want to check
	// first if we've computed this argument!
	result := libgollatz.Collatz(argument)

	// This is the first step toward caching results: we want to switch on
	// r.Method. Normally, it's not good to put "for now" in comments,
	// but, for now, we're only going to check for a GET request.
	switch r.Method {
	case "GET":
		// This was broken out into a function, but as we are currently only
		// serving to a browser, there's no reason to break it out. This is
		// subject to change once we implement a proper REST API. Then, we
		// can execute this based on seeing a browser user agent.
		err = templates.ExecuteTemplate(w, "result.html", result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	default:
		http.Error(w, "Only GET requests allowed right now. Sorry!", http.StatusMethodNotAllowed)
	}

}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func main() {
	http.HandleFunc("/collatz/", collatzHandler)
	http.HandleFunc("/", rootHandler)

	http.ListenAndServe(":8080", nil)
}
