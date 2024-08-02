package artgen

import (
	"fmt"
	"strings"
)

// Reads input text,gets the pattern convert it to ascii art
func PrintingAscii(text, patternFile string) (string, error) {
	res := ""
	asciiMap, err := AsciiMapping(patternFile)
	if err != nil {
		return "", err
	}

	for n := 0; n < 8; n++ {
		for _, ch := range text {
			res += asciiMap[ch][n]
		}
		res += "\n"
	}

	return res, nil
}

func HandleInput(text string) (string, error) {
	//text = strings.ReplaceAll(text, "\n", "\\n")
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
