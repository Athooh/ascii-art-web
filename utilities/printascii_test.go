package utils

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func Test_PrintAsciiArt(t *testing.T) {
	// Compute the correct path to the standard.txt file
	basePath, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	asciiFilePath := filepath.Join(basePath, "..", "banners/standard.txt")

	// Load the ASCII characters from the sample banner file for testing
	asciiChars, err := LoadAsciiChars(asciiFilePath)
	if err != nil {
		t.Fatalf("Failed to load ASCII characters: %v", err)
	}

	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:  "Simple Text",
			input: "Hello",
			output: ` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := CaptureOutput(func() {
				PrintAsciiArt(tt.input, asciiChars)
			})

			if output != tt.output {
				t.Errorf("Expected output:\n%s\nGot:\n%s", tt.output, output)
			}
		})
	}
}

// Helper function to capture stdout
func CaptureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
