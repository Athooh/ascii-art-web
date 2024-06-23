package utils

import (
	"bufio"
	"fmt"
	"os"
)

const (
	expectedAsciiChars = 95
	expectedlines = 8
)

// LoadAsciiChars loads ASCII characters from a file into a map.
func LoadAsciiChars(filename string) (map[byte][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("file '%s' not found", filename)
		} else {
			return nil, fmt.Errorf("opening file: %w", err)
		}
	}
	defer file.Close()

	asciiChars := make(map[byte][]string)

	scanner := bufio.NewScanner(file)

	currentChar := byte(' ') // or 32
	count := 0
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		// Append lines to current character's value slice until 8 lines are read.
		if count != expectedlines {
			asciiChars[currentChar] = append(asciiChars[currentChar], line)
			count++
		} else {
			currentChar++
			count = 0
		}
	}

	if len(asciiChars) == 0 {
		return nil, fmt.Errorf("file '%s' is empty", filename)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	// Print error for incorrect number of ASCII characters
	if len(asciiChars) != expectedAsciiChars {
		return nil, fmt.Errorf("expected %d but got %d ASCII CHARACTERS. Ensure you use the correct number of ASCII CHARACTERS", expectedAsciiChars, len(asciiChars))
	}

	// Print error for incorrect number of lines per character
	if count != expectedlines {
		return nil, fmt.Errorf("expected %d lines but got %d lines. Ensure the file has the correct number of lines per character", expectedlines, count)
	}

	return asciiChars, nil
}
