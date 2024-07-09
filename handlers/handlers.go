package handler

import (
	"html/template"
	"log"
	"net/http"
	utils "web/utilities"
)

type PageData struct {
	Text  string
	Art   string
	Error string
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "HTTP status 404 - page not found", http.StatusNotFound)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/form.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "HTTP status 500 - Internal Server Errors", http.StatusInternalServerError)
	}
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("AsciiArtHandler called")
	if r.Method != http.MethodPost {
		log.Println("Method not allowed")
		http.Error(w, "HTTP status 405 - method not allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")
	log.Printf("Text: %s, Banner: %s", text, banner)
	pageData := PageData{Text: text}

	if text == "" || containsNonASCII(text) {
		pageData.Error = "HTTP status 400 - Bad Request: 'text' parameter is required and should contain only ASCII characters."
		renderForm(w, pageData, r)
		return
	}

	if banner == "" {
		banner = "standard"
	}

	asciiChars, err := utils.LoadAsciiChars("banners/" + banner + ".txt")
	if err != nil {
		log.Printf("Error loading banner: %v", err)
		pageData.Error = "HTTP status 500 internal Server Error - could not load banner"
		renderForm(w, pageData, r)
		return
	}

	art, err := utils.GenerateAsciiArt(text, asciiChars)
	if err != nil {
		log.Printf("Error generating ASCII art: %v", err)
		pageData.Error = err.Error()
		renderForm(w, pageData, r)
		return
	}

	pageData.Art = art
	renderForm(w, pageData, r)
}

func containsNonASCII(text string) bool {
	for _, char := range text {
		if char > 127 {
			return true
		}
	}
	return false
}

func renderForm(w http.ResponseWriter, data PageData, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/form.html"))
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error rendering form: %v", err)
		http.Error(w, "HTTP status 500 - Internal Server Error", http.StatusInternalServerError)
	}
}
