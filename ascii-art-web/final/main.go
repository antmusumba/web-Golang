package main

import (
	"fmt"
	"log"
	"net/http"
	"output/utils"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", utils.GetAsciiForm)
	http.HandleFunc("/ascii-art", utils.PostAsciiArt)
	fmt.Println("SUCCESS!! listen to server at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
