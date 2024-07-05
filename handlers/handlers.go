package handler

import (
	"html/template"
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
		http.Error(w, "HTTP status 500 - Internal Server Errors", http.StatusInternalServerError)
	}
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	pageData := PageData{Text: text}

	if r.Method != http.MethodPost {
		http.Error(w, "HTTP status 405 - method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if text == "" {
		pageData.Error = "Bad Request: 'text' parameter is required"
		renderForm(w, pageData, r)
		return
	}

	if banner == "" {
		banner = "standard"
	}

	// Check for non-ASCII characters
	if containsNonASCII(text) {
		pageData.Error = "Input contains non-ASCII characters or emojis. Please use only ASCII characters."
		renderForm(w, pageData, r)
		return
	}

	// Load ASCII characters from the specified file in the 'banners' directory
	asciiChars, err := utils.LoadAsciiChars("banners/" + banner + ".txt")
	if err != nil {
		pageData.Error = "500 internal server error: could not load banner"
		renderForm(w, pageData, r)
		return
	}

	// Generate ASCII art
	art, err := utils.GenerateAsciiArt(text, asciiChars)
	if err != nil {
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
	if r.Header.Get("X-Requested-With") == "XMLHttpRequest" {
		if data.Error != "" {
			w.Write([]byte(data.Error))
		} else {
			w.Write([]byte(data.Art))
		}
	} else {
		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
			return
		}
	}
}
