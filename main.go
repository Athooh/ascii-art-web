package main

import (
	"fmt"
	"net/http"

	handler "web/handlers"
)

func main() {
	http.HandleFunc("/", handler.FormHandler)
	http.HandleFunc("/ascii-art", handler.AsciiArtHandler)
	http.HandleFunc("/favicon.ico", http.NotFound) // Handle favicon.ico requests
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	fmt.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
