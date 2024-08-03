package main

import (
    "html/template"
    "log"
    "net/http"
)

type PageData struct {
    Title   string
    Content string
}

func homePage(w http.ResponseWriter, r *http.Request) {
    data := PageData{
        Title:   "Home Page",
        Content: "Welcome to the Home Page!",
    }
    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
        return
    }
    err = tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
    }
}

func main() {
    http.HandleFunc("/", homePage)
    log.Println("Server started at :8081")
    err := http.ListenAndServe(":8081", nil)
    if err != nil {
        log.Fatalf("Could not start server: %s\n", err.Error())
    }
}
