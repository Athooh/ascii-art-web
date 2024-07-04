package utils

import (
	"fmt"
	"strings"
)

// PrintAsciiArt prints the given text as ASCII art using the provided map of characters.
func PrintAsciiArt(text string, asciiChars map[byte][]string) {
	// Check if any character is outside the ASCII range (32-127)
	for _, char := range text {
		if char > 127 || char < 32 {
			fmt.Printf("Error: Character %q is not accepted\n", char)
			return
		}
	}

	// Print each line of the ASCII art
	for i := 0; i < 8; i++ {
		PrintLine(text, asciiChars, i)
		fmt.Println()
	}
}

// PrintLine prints a single line of the ASCII art for the given text.
func PrintLine(text string, asciiChars map[byte][]string, line int) {
	for _, char := range text {
		fmt.Print(asciiChars[byte(char)][line]) // Print the ASCII representation of the character
	}
}

func GenerateAsciiArt(text string, asciiChars map[byte][]string) string {
	var result strings.Builder

	for _, line := range strings.Split(text, "\n") {
		for i := 0; i < 8; i++ {
			for _, char := range line {
				result.WriteString(asciiChars[byte(char)][i])
			}
			result.WriteString("\n")
		}
		result.WriteString("\n")
	}
	return result.String()
}
