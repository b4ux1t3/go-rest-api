package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/b4ux1t3/libgollatz"
)

// RequestNumber doesn't need to exist, I'm just having issues figuring something out
// Update from future me: So, the problem is that the function ExecuteTemplate is
// expecting a type with a data interface ,which a primative does not fulfill. There is
// no good way around this as of right now, except to hardcode the input number into
// the template.
//TODO: Get rid of this
type RequestNumber struct {
	Num int
}

var templates = template.Must(template.ParseFiles("result.html", "index.html"))

func collatzHandler(w http.ResponseWriter, r *http.Request) {
	argumentIndex := len("/collatz/")
	argument, err := strconv.ParseUint(r.URL.Path[argumentIndex:], 10, 64)
	// This executes if we do not receive aninteger from the user after "/collatz/" in the URL path.
	if err != nil {
		//http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		http.Error(w, "Unable to parse argument. URL path should be: /collatz/<POSITIVE BASE 10 INTEGER>", http.StatusInternalServerError)
		return
	}

	result := libgollatz.Collatz(argument)

	// This was broken out into a function, but as we are currently only
	// serving to a browser, there's no reason to break it out. This is
	// subject to change once we implement a proper REST API. Then, we
	// can execute this based on seeing a browser user agent.
	err = templates.ExecuteTemplate(w, "result.html", result)
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
