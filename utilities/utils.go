package utils

import (
	"strings"

	asciiArt "web/ascii-funcs"
)

func GenerateAsciiArt(text string, asciiChars map[byte][]string) string {
	var result strings.Builder
	text = asciiArt.ReplaceSpecChars(text)

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
