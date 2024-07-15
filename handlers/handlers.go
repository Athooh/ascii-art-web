package handler

import (
	"html/template"
	"log"
	"net/http"

	utils "web/utilities" // Import utilities package for helper functions
)

// PageData struct holds data to be rendered in templates
type PageData struct {
	Text  string // Input text from the form
	Art   string // Generated ASCII art based on input text and banner
	Error string // Error message, if any
}

// FormHandler handles GET requests to render the form template
func FormHandler(w http.ResponseWriter, r *http.Request) {
	// Handle requests only at root URL
	if r.URL.Path != "/" {
		http.Error(w, "HTTP status 404 - Page not found", http.StatusNotFound)
		return
	}

	// Parse and execute form.html template
	tmpl := template.Must(template.ParseFiles("templates/form.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "HTTP status 500 - Internal Server Error", http.StatusInternalServerError)
	}
}

// AsciiArtHandler handles POST requests to generate ASCII art based on form input
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("AsciiArtHandler called")

	// Check if request method is POST
	if r.Method != http.MethodPost {
		log.Println("Method not allowed")
		http.Error(w, "HTTP status 405 - Method not allowed", http.StatusMethodNotAllowed)
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
		pageData.Error = "HTTP status 400 - Bad Request: 'text' parameter is required and should contain only ASCII characters."
		renderForm(w, pageData, r)
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
		pageData.Error = "HTTP status 500 - Internal Server Error: could not load banner"
		renderForm(w, pageData, r)
		return
	}

	// Generate ASCII art based on text and ASCII characters
	art, err := utils.GenerateAsciiArt(text, asciiChars)
	if err != nil {
		log.Printf("Error generating ASCII art: %v", err)
		pageData.Error = err.Error()
		renderForm(w, pageData, r)
		return
	}

	// Assign generated ASCII art to PageData
	pageData.Art = art

	// Render form template with updated PageData
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

// renderForm renders the form template with provided data
func renderForm(w http.ResponseWriter, data PageData, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/form.html"))
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error rendering form: %v", err)
		http.Error(w, "HTTP status 500 - Internal Server Error", http.StatusInternalServerError)
	}
}
