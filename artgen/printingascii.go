package artgen

import (
	"fmt"
	"strings"
)

// Reads input text,gets the pattern convert it to ascii art
func PrintingAscii(text, patternFile string) (string, error) {
	res := ""
	lines := strings.Split(text, "\n")
	asciiMap, err := AsciiMapping(patternFile)
	if err != nil {
		return "", err
	}

	for _, word := range lines { // case of multiple newlines

		for n := 0; n < 8; n++ {
			for _, ch := range word {
				res += asciiMap[ch][n]
			}
			res += "\n"
		}
	}

	return res, nil
}

// HandleInput ensures that the text contains only printable characters
func HandleInput(text string) (string, error) {
	text = strings.ReplaceAll(text, "\r", "")

	for _, char := range text {
		if char < 32 || char > 126 {
			if char == 10 || char == 13 {
				continue
			} else {
				return "", fmt.Errorf(" char is not valid")
			}
		}
	}
	return text, nil
}
