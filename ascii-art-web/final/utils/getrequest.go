package utils

import (
	"html/template"
	"log"
	"net/http"
)

func GetAsciiForm(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		if r.Method != http.MethodGet {
			http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		t, err := template.ParseFiles("template/index.html")
		if err != nil {
			log.Printf("Error parsing template: %v\n", err)
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		err = t.Execute(w, nil)
		if err != nil {
			log.Printf("Error executing template: %v\n", err)
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "400 path not found", http.StatusBadRequest)
		return
	}
}
