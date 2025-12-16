package main

import (
	"os"
	"strings"
)

const (
	height    = 8
	firstRune = 32  // space
	lastRune  = 126 // ~
	runeCount = lastRune - firstRune + 1
)

func LoadBanner(path string) (map[rune][]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Normalize CRLF to LF to avoid injecting carriage returns into the output
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	content = strings.ReplaceAll(content, "\r", "\n")
	lines := strings.Split(content, "\n")
	banner := make(map[rune][]string)

	// Each rune uses 9 lines in file: 1 empty + 8 drawing lines
	for i := 0; i < runeCount; i++ {
		start := i*(height+1) + 1
		end := start + height
		if end > len(lines) {
			break
		}
		r := rune(firstRune + i)
		banner[r] = lines[start:end]
	}
	return banner, nil
}

func RenderASCII(text string, banner map[rune][]string) string {
	// Support both:
	// - literal "\n" (like your CLI project)
	// - real newlines from textarea
	text = strings.ReplaceAll(text, "\r\n", "\n")
	text = strings.ReplaceAll(text, `\n`, "\n")

	inputLines := strings.Split(text, "\n")

	var out strings.Builder

	for i, line := range inputLines {
		if line == "" {
			if i != len(inputLines)-1 {
				out.WriteString("\n")
			}
			continue
		}

		for row := 0; row < height; row++ {
			for _, r := range line {
				if r < firstRune || r > lastRune {
					continue
				}
				block, ok := banner[r]
				if !ok || len(block) <= row {
					continue
				}
				out.WriteString(block[row])
			}
			out.WriteString("\n")
		}
	}

	return out.String()
}
