package wc

import (
	"bytes"
	"strings"
)

func CountBytes(content []byte) int {
	return len(content)
}

func CountLines(content []byte) int {
	if len(content) == 0 {
		return 0
	}

	return 1 + strings.Count(string(content), "\n")
}

func CountWords(content []byte) int {
	return len(bytes.Fields(content))
}

func CountCharacters(content []byte) int {
	return len(bytes.Runes(content))
}
