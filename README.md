# ASCII Art Web 

## Table of Contents

- [Introduction](#introduction)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
  - [Main Function](#main-function)
  - [Route Handlers](#route-handlers)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)
- [Acknowledgments](#acknowledgments)
- [Implementation](#implementation)

## Introduction

The ASCII Art Web project is a simple Go-based web server that handles different routes and serves static files. It provides a user interface for generating ASCII art from text input and aims to make the site more appealing, interactive, intuitive, and user-friendly.

## Getting Started

Follow these instructions to get the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (>= 1.16)
- Web browser

### Installation

1. Clone the repository:
   ```sh
   git clone https://learn.zone01kisumu.ke/git/seodhiambo/ascii-art-web.git
   cd ascii-art-web
   ```

2. Run the server:
   ```sh
   go run .
   ```

## Usage

1. Navigate to http://localhost:8000 in your web browser.
2. Enter the text you want to convert to ASCII art in the form.
3. Select a banner style (optional).
4. Submit the form to receive the generated ASCII art.

### Main Function

The `main` function sets up the server and routes:

1. **Root Handler**: Maps the root URL `("/")` to the `FormHandler` function in the `handler` package.
2. **ASCII Art Handler**: Maps the `/ascii-art` URL to the `AsciiArtHandler` function in the `handler` package.
3. **Favicon Handler**: Maps requests for the favicon (`/favicon.ico`) to the `http.NotFound` handler, which returns a 404 not found response.
4. **Static File Server**: Serves static files from the `./static` directory. Requests to URLs starting with `/static/` will have the `/static/` prefix stripped before looking for the corresponding file in the `./static` directory.

The server starts on port 8000 and handles any startup errors.

### Route Handlers

The `handlers.go` file contains the implementation of the handlers referenced in the `main.go` file. These handlers manage specific routes and perform various tasks based on incoming HTTP requests.

#### Handlers Overview

1. **Form Handler**: Manages the root URL (/) and likely serves a form for user input.
2. **ASCII Art Handler**: Processes the `/ascii-art` URL and generates ASCII art from the input text.
3. **Favicon Handler**: Manages requests for the favicon (`/favicon.ico`), returning a 404 not found response.

## Configuration

### Environment Variables

Configure the following environment variables if necessary:
- `PORT`: The port number the server will listen on (default: 8000).

## Contributing

1. Fork the repository
2. Create a new branch (`git checkout -b feature/YourFeature`)
3. Commit your changes (`git commit -m 'Add some feature'`)
4. Push to the branch (`git push origin feature/YourFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See [LICENSE](./LICENSE) for more information.

## Contact

For any inquiries, please contact the project maintainers:

- [Stella Oiro](https://learn.zone01kisumu.ke/git/steoiro)
- [Seth Odhiambo](https://learn.zone01kisumu.ke/git/seodhiambo)

## Acknowledgments

- [Go Documentation](https://golang.org/doc/)
- [Standard Library Packages](https://pkg.go.dev/std)

## Implementation

### Description

The ASCII Art Web project converts user input text into ASCII art using a Go-based web server. The site is designed to be more appealing, interactive, intuitive, and user-friendly by incorporating CSS for better design and user feedback.

### Authors

- Stella Oiro [@steoiro](https://learn.zone01kisumu.ke/git/steoiro)
- Seth Odhiambo [@seodhiambo](https://learn.zone01kisumu.ke/git/seodhiambo)

### Usage: How to Run

1. Clone the repository:
   ```sh
   git clone https://learn.zone01kisumu.ke/git/seodhiambo/ascii-art-web.git
   cd ascii-art-web
   ```

2. Run the server:
   ```sh
   go run .
   ```

3. Open your web browser and navigate to http://localhost:8000.

### Implementation Details: Algorithm

- **Package Declaration**:
  ```go
  package main
  ```

- **Imports**:
  ```go
  import (
      "fmt"
      "net/http"
      handler "web/handlers"
  )
  ```

- **Main Function**:
  - Maps the root URL `("/")` to the `FormHandler` function.
  - Maps the `/ascii-art` URL to the `AsciiArtHandler` function.
  - Maps requests for the favicon (`/favicon.ico`) to the `http.NotFound` handler.
  - Serves static files from the `./static` directory.

- **Route Handlers**:
  - **Form Handler**: Manages the root URL and serves a form for user input.
  - **ASCII Art Handler**: Processes the `/ascii-art` URL and generates ASCII art from the input text.
  - **Favicon Handler**: Manages requests for the favicon, returning a 404 not found response.