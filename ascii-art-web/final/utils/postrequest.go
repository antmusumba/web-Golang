package utils

import (
    "html/template"
    "log"
    "net/http"
    "os"
    "strings"
)

type Data struct {
    Filename string
    Input    string
    Result   string
}

func PostAsciiArt(w http.ResponseWriter, r *http.Request) {
    log.Printf("Received %s request on /ascii-art route\n", r.Method)

    if r.Method != http.MethodPost {
        http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    action := r.FormValue("action")
    if action == "Reset" {
        t, err := template.ParseFiles("template/index.html")
        if err != nil {
            log.Printf("Error parsing template: %v\n", err)
            http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
            return
        }
        err = t.Execute(w, &Data{})
        if err != nil {
            log.Printf("Error executing template: %v\n", err)
            http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
            return
        }
        return
    }

    text := r.FormValue("word")
    banner := strings.TrimSpace(r.FormValue("banner"))
    log.Printf("Received banner value: %s\n", banner)

    switch banner {
    case "standard":
        banner = "banners/standard.txt"
    case "thinkertoy":
        banner = "banners/thinkertoy.txt"
    case "shadow":
        banner = "banners/shadow.txt"
    default:
        http.Error(w, "400 Bad Request: Invalid banner", http.StatusBadRequest)
        return
    }

    if text == "" || banner == "" {
        http.Error(w, "400 Bad Request: Missing text or banner", http.StatusBadRequest)
        return
    }

    bannerFileContent, err := os.ReadFile(banner)
    if err != nil {
        log.Printf("Error reading banner file: %v\n", err)
        http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
        return
    }

    log.Printf("Banner file content length: %d\n", len(bannerFileContent))
    bannerSlice := SplitFile(string(bannerFileContent), banner)
    if len(bannerSlice) != 856 {
        http.Error(w, "500 Internal Server Error: Banner content is altered", http.StatusInternalServerError)
        return
    }

    // s := ReplaceEscape(text)
    // for _, char := range s {
    //     if char > 126 || char < 32 {
    //         http.Error(w, "400 Bad Request: Invalid character in input", http.StatusBadRequest)
    //         return
    //     }
    // }

    result := DisplayText(text, bannerSlice)

    resultData := &Data{
        Filename: banner,
        Input:    text,
        Result:   result,
    }

    t, err := template.ParseFiles("template/index.html")
    if err != nil {
        log.Printf("Error parsing template: %v\n", err)
        http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
        return
    }

    err = t.Execute(w, resultData)
    if err != nil {
        log.Printf("Error executing template: %v\n", err)
        http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
        return
    }
}
