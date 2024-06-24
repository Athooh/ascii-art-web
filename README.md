# ASCII Art Web Server

This repository contains a simple Go web server that allows users to generate ASCII art from text. The server handles different routes, serves static files, and uses HTML templates for rendering pages.

## Table of Contents

1. [Overview](#overview)
2. [Features](#features)
3. [Installation](#installation)
4. [Usage](#usage)
5. [Routes](#routes)
6. [Handlers](#handlers)
7. [Static Files](#static-files)
8. [Error Handling](#error-handling)
9. [License](#license)

## Overview

The ASCII Art Web Server is a Go application that provides a web interface for generating ASCII art from text inputs. Users can submit text through a web form, choose a banner style, and receive the corresponding ASCII art as output.

## Features

- Serve HTML pages using templates.
- Handle various routes, including form submission and static file serving.
- Generate ASCII art from text inputs.
- Handle different HTTP methods and provide appropriate responses.

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/ascii-art-web-server.git
    cd ascii-art-web-server
    ```

2. Build the Go application:
    ```bash
    go build
    ```

3. Run the server:
    ```bash
    ./ascii-art-web-server
    ```

The server will start on `http://localhost:8080`.

## Usage

1. Navigate to `http://localhost:8080` in your web browser.
2. Enter the text you want to convert to ASCII art in the form.
3. Select a banner style (optional).
4. Submit the form to receive the generated ASCII art.

## Routes

- `/`: Root URL, serves the form for ASCII art generation.
- `/ascii-art`: Handles form submissions and generates ASCII art.
- `/favicon.ico`: Handles requests for the favicon, returns a 404 not found response.
- `/static/`: Serves static files from the `./static` directory.

## Handlers

### FormHandler

Handles requests to the root URL (`/`). Serves the form for ASCII art generation.

```go
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
```

### AsciiArtHandler

Handles form submissions (`/ascii-art`). Generates ASCII art from the provided text and banner.

```go
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

    asciiChars, err := asciiArt.LoadAsciiChars("banners/" + banner + ".txt")
    if err != nil {
        http.Error(w, "500 internal server error: could not load banner", http.StatusInternalServerError)
        return
    }

    art := utils.GenerateAsciiArt(text, asciiChars)

    tmpl := template.Must(template.ParseFiles("templates/result.html"))
    err = tmpl.Execute(w, art)
    if err != nil {
        http.Error(w, "500 internal server error", http.StatusInternalServerError)
    }
}
```

## Static Files

Static files are served from the `./static` directory. Requests to URLs starting with `/static/` will have the `/static/` prefix stripped before looking for the corresponding file in the `./static` directory.

```go
http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
```

## Error Handling

The server includes basic error handling for various scenarios, such as:
- Returning a 404 error for invalid URLs.
- Returning a 405 error for unsupported HTTP methods.
- Handling a 500 - internal server error during template execution or file loading.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

