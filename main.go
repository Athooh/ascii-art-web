package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	asciiArt "web/ascii-funcs"
)

func main() {
	http.HandleFunc("/", formHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("form.html"))
	tmpl.Execute(w, nil)
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if banner == "" {
		banner = "standard"
	}

	// Load ASCII characters from the specified file
	asciiChars, err := asciiArt.LoadAsciiChars("banners/" + banner + ".txt")
	if err != nil {
		http.Error(w, "Error loading banner. Did you mean 'thinkertoy', 'shadow', or 'standard'?", http.StatusInternalServerError)
		return
	}

	// Generate ASCII art
	art := generateAsciiArt(text, asciiChars)

	tmpl := template.Must(template.ParseFiles("result.html"))
	tmpl.Execute(w, art)
}

func generateAsciiArt(text string, asciiChars map[byte][]string) string {
	var result strings.Builder
	text = asciiArt.ReplaceSpecChars(text)

	for i := 0; i < 8; i++ {
		for _, char := range text {
			if char == '\n' {
				result.WriteString("\n")
			} else {
				result.WriteString(asciiChars[byte(char)][i])
			}
		}
		result.WriteString("\n")
	}

	return result.String()
}
