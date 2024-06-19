package handler

import (
	"html/template"
	"net/http"

	asciiArt "web/ascii-funcs"
	utils "web/utilities"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "HTTP status 404 - page not found", http.StatusNotFound)
		return
	}

	tmpl := template.Must(template.ParseFiles("form.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "HTTP status 500 - Internal Server Errors", http.StatusInternalServerError)
	}
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "HTTP status 405 - method not allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" {
		http.Error(w, "HTTP status 400 - Bad Request: 'text' parameter is required", http.StatusBadRequest)
		return
	}

	if banner == "" {
		banner = "standard"
	}

	// Load ASCII characters from the specified file in the 'banners' directory
	asciiChars, err := asciiArt.LoadAsciiChars("banners/" + banner + ".txt")
	if err != nil {
		http.Error(w, "500 internal server error: could not load banner", http.StatusInternalServerError)
		return
	}

	// Generate ASCII art
	art := utils.GenerateAsciiArt(text, asciiChars)

	tmpl := template.Must(template.ParseFiles("result.html"))
	err = tmpl.Execute(w, art)
	if err != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
	}
}
