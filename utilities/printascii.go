package utils

import (
	"fmt"
	"strings"
)

// PrintAsciiArt prints the given text as ASCII art using the provided map of characters.
func PrintAsciiArt(text string, asciiChars map[byte][]string) error {
	for _, char := range text {
		if char > 127 || char < 32 {
			return fmt.Errorf("character %q is not accepted", char)
		}
	}

	for i := 0; i < 8; i++ {
		PrintLine(text, asciiChars, i)
		fmt.Println()
	}
	return nil
}

// PrintLine prints a single line of the ASCII art for the given text.
func PrintLine(text string, asciiChars map[byte][]string, line int) {
	for _, char := range text {
		fmt.Print(asciiChars[byte(char)][line]) // Print the ASCII representation of the character
	}
}

func GenerateAsciiArt(text string, asciiChars map[byte][]string) (string, error) {
	var result strings.Builder

	for _, line := range strings.Split(text, "\n") {
		for _, char := range line {
			if char > 127 || char < 32 {
				return "", fmt.Errorf("character %q is not accepted", char)
			}
		}

		for i := 0; i < 8; i++ {
			for _, char := range line {
				result.WriteString(asciiChars[byte(char)][i])
			}
			result.WriteString("\n")
		}
		result.WriteString("\n")
	}
	return result.String(), nil
}
