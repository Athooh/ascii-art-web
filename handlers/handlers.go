package handler

import (
	"html/template"
	"log"
	"net/http"
	utils "web/utilities"
)

type PageData struct {
	Text    string
	Art     string
	Error   string
	Code    int
	Message string
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		renderError(w, 404, "HTTP status 404 - Page not found")
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/form.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		renderError(w, 500, "HTTP status 500 - Internal Server Error")
	}
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("AsciiArtHandler called")
	if r.Method != http.MethodPost {
		log.Println("Method not allowed")
		renderError(w, 405, "HTTP status 405 - Method not allowed")
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")
	log.Printf("Text: %s, Banner: %s", text, banner)
	pageData := PageData{Text: text}

	if text == "" || containsNonASCII(text) {
		renderError(w, 400, "HTTP status 400 - Bad Request")
		return
	}

	if banner == "" {
		banner = "standard"
	}

	asciiChars, err := utils.LoadAsciiChars("banners/" + banner + ".txt")
	if err != nil {
		log.Printf("Error loading banner: %v", err)
		renderError(w, 500, "HTTP status 500 - Internal Server Error: Could not load banner")
		return
	}

	art, err := utils.GenerateAsciiArt(text, asciiChars)
	if err != nil {
		log.Printf("Error generating ASCII art: %v", err)
		renderError(w, 500, err.Error())
		return
	}

	pageData.Art = art
	renderForm(w, pageData)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/about.html"))
	tmpl.Execute(w, nil)
}

func HowItWorksHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/how-it-works.html"))
	tmpl.Execute(w, nil)
}

func containsNonASCII(text string) bool {
	for _, char := range text {
		if char > 127 {
			return true
		}
	}
	return false
}

func renderForm(w http.ResponseWriter, data PageData) {
	tmpl := template.Must(template.ParseFiles("templates/form.html"))
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error rendering form: %v", err)
		renderError(w, 500, "HTTP status 500 - Internal Server Error")
	}
}

func renderError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	tmpl := template.Must(template.ParseFiles("templates/error.html"))
	err := tmpl.Execute(w, PageData{Code: code, Message: message})
	if err != nil {
		log.Printf("Error rendering error page: %v", err)
		http.Error(w, "HTTP status 500 - Internal Server Error", http.StatusInternalServerError)
	}
}
