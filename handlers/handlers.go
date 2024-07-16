package handler

import (
	"html/template"
	"log"
	"net/http"

	utils "web/utilities" // Import utilities package for helper functions
)

// PageData struct holds data to be rendered in templates
type PageData struct {
	Text  string
	Art   string
	Error string
}

// FormHandler handles GET requests to render the form template
func FormHandler(w http.ResponseWriter, r *http.Request) {
	// Handle requests only at root URL
	if r.URL.Path != "/" {
		http.Error(w, "HTTP status 404 - page not found", http.StatusNotFound)
		return
	}

	// Parse and execute form.html template
	tmpl := template.Must(template.ParseFiles("templates/form.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "HTTP status 500 - Internal Server Errors", http.StatusInternalServerError)
	}
}

// AsciiArtHandler handles POST requests to generate ASCII art based on form input
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("AsciiArtHandler called")

	// Check if request method is POST
	if r.Method != http.MethodPost {
		log.Println("Method not allowed")
		http.Error(w, "HTTP status 405 - method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Retrieve text and banner values from form
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	log.Printf("Text: %s, Banner: %s", text, banner)

	// Initialize PageData struct with input text
	pageData := PageData{Text: text}

	// Validate text input for non-ASCII characters
	if text == "" || containsNonASCII(text) {
		renderError(w, 400, "HTTP status 400 - Bad Request")
		return
	}

	// Default to 'standard' banner if none specified
	if banner == "" {
		banner = "standard"
	}

	// Load ASCII characters from banner file
	asciiChars, err := utils.LoadAsciiChars("banners/" + banner + ".txt")
	if err != nil {
		log.Printf("Error loading banner: %v", err)
		pageData.Error = "HTTP status 500 internal Server Error - could not load banner"
		renderForm(w, pageData, r)
		return
	}

	// Generate ASCII art based on text and ASCII characters
	art, err := utils.GenerateAsciiArt(text, asciiChars)
	if err != nil {
		log.Printf("Error generating ASCII art: %v", err)
		renderError(w, 500, err.Error())
		return
	}

	// Assign generated ASCII art to PageData
	pageData.Art = art
	renderForm(w, pageData, r)
}

// containsNonASCII checks if a string contains non-ASCII characters
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
