package asciiArt

import (
	"fmt"
	"strings"
)

// ReplaceSpecChars replaces escape sequences with their corresponding special characters.
func ReplaceSpecChars(s string) string {
	replace := strings.NewReplacer(
		"\\r", "\r", // Replace '\r' with carriage return
		"\\b", "\b", // Replace '\b' with backspace
		"\\t", "    ", // Replace '\t' with four spaces
		"\\f", "\f", // Replace '\f' with form feed
		"\\a", "\a", // Replace '\a' with alert (bell)
		"\\v", "\v", // Replace '\v' with vertical tab
	)
	return replace.Replace(s)
}

// PrintAsciiArt prints the given text as ASCII art using the provided map of characters.
func PrintAsciiArt(text string, asciiChars map[byte][]string) {
	text = ReplaceSpecChars(text)
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
		if char == '\n' {
			fmt.Println()
		} else {
			fmt.Print(asciiChars[byte(char)][line]) // Print the ASCII representation of the character
		}
	}
}

// ProcessArguments processes the input arguments and prints ASCII art for each argument.
func ProcessArguments(args string, asciiChars map[byte][]string) {
	arguments := strings.Split(args, "\\n") // Split arguments by '\n'
	// Handle empty input or newlines
	countSpaces := 0
	for _, arg := range arguments {
		if arg == "" {
			countSpaces++
			if countSpaces < len(arguments) {
				fmt.Println() // Print a newline for each empty argument except the last one
			}
		} else {
			PrintAsciiArt(arg, asciiChars) // Print ASCII art for the non-empty argument
		}
	}
}
