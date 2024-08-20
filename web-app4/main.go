package main

import (
	"fmt"
	"net/http"
)

func handleForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "unable to parse form", http.StatusBadRequest)
			return
		}

		// Extract form data
		name := r.FormValue("name")
		email := r.FormValue("email")
		message := r.FormValue("message")

		// Process the form data (e.g., save to a database or send an email)
		fmt.Fprintf(w, "Name: %s\nEmail: %s\nMessage: %s", name, email, message)
	} else {
	}
}
func main() {
    http.HandleFunc("/submit_form", handleForm)
    http.ListenAndServe(":8080", nil)
}