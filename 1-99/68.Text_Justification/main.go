package main

import (
	"fmt"
)

func fullJustify(words []string, maxWidth int) []string {
	res := make([]string, 0)
	for len(words) > 0 {
		size := wordSizeOfLine(words, maxWidth)
		var line string
		if size == len(words) {
			line = formatLastLine(size, maxWidth, words)
		} else {
			line = formatLine(size, maxWidth, words)
		}

		res = append(res, line)
		words = words[size:]
	}
	return res
}

func wordSizeOfLine(words []string, maxWidth int) int {
	length := 0
	i := 0
	for i < len(words) && length+len(words[i])+i <= maxWidth {
		length += len(words[i])
		i++
	}
	return i
}

func formatLine(size, maxWidth int, words []string) string {
	length := 0
	for i := 0; i < size; i++ {
		length += len(words[i])
	}

	spaceLen := maxWidth - length
	extraLeft := 0
	if size > 1 {
		extraLeft = spaceLen % (size - 1)
		spaceLen = spaceLen / (size - 1)
	}

	line := ""
	for i := 0; i < size; i++ {
		line += words[i]

		if len(line) < maxWidth {
			line = addSpace(line, spaceLen)
			if extraLeft > 0 {
				line = addSpace(line, 1)
				extraLeft--
			}
		}
	}
	return line
}

func formatLastLine(size, maxWidth int, words []string) string {
	line := ""
	for i, word := range words {
		line += word
		if i < len(words)-1 {
			line += " "
		}
	}

	for len(line) < maxWidth {
		line += " "
	}
	return line
}

func addSpace(s string, spaceLen int) string {
	for i := 0; i < spaceLen; i++ {
		s += " "
	}
	return s
}

func main() {
	words := []string{"ask", "not", "what", "your", "country", "can", "do", "for", "you", "ask", "what", "you", "can", "do", "for", "your", "country"}
	maxWidth := 16
	for _, s := range fullJustify(words, maxWidth) {
		fmt.Println(s, len(s))
	}
}
