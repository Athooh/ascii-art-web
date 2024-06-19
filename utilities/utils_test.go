package utils

import "testing"

func TestGenerateAsciiArt(t *testing.T) {
	// Sample ASCII art characters for 'A' and 'B'
	asciiChars := map[byte][]string{}

	tests := []struct {
		name string
		args struct {
			text       string
			asciiChars map[byte][]string
		}
		want string
	}{
		{
			name: "Single character A",
			args: struct {
				text       string
				asciiChars map[byte][]string
			}{
				text:       "A",
				asciiChars: asciiChars,
			},
			want: "                /\\         /  \\       / /\\ \\     / ____ \\   /_/    \\_\\                         ",
		},
		{
			name: "Single character B",
			args: struct {
				text       string
				asciiChars map[byte][]string
			}{
				text:       "B",
				asciiChars: asciiChars,
			},
			want: " ____    |  _ \\   | |_) |  |  _ <   | |_) |  |____/                    ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateAsciiArt(tt.args.text, tt.args.asciiChars); got != tt.want {
				t.Errorf("GenerateAsciiArt() = %v, want %v", got, tt.want)
			}
		})
	}
}
