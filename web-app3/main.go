package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
)

// Struct to hold form data
type FormData struct {
    Username string
    Password string
}

// Handler to render the form
func formHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/form.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, nil)
}

// Handler to process form submission
func loginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    username := r.FormValue("username")
    password := r.FormValue("password")

    formData := FormData{
        Username: username,
        Password: password,
    }

    // You can now use username and password to authenticate the user
    fmt.Fprintf(w, "Username: %s, Password: %s", formData.Username, formData.Password)
}

func main() {
    http.HandleFunc("/", formHandler)
    http.HandleFunc("/submit", loginHandler)

    fmt.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
