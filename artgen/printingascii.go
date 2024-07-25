package artgen

import (
	"fmt"
	"strings"
)

// Reads input text,gets the pattern convert it to ascii art
func PrintingAscii(text, patternFile string) (string, error) {
	text = strings.ReplaceAll(text, "\n", "\\n")
	res := ""

	for _, char := range text {
		if char < 32 || char > 126 {
			return "", fmt.Errorf(" char is not valid")
		}
	}

	lines := strings.Split(text, "\\n")
	asciiMap, err := AsciiMapping(patternFile)
	if err != nil {
		return "", err
	}

	count := 0

	for _, word := range lines { // case of multiple newlines
		if word == "" {
			count++
			if count < len(lines) {
				res += "\n"
			}
		} else {
			for n := 0; n < 8; n++ {
				for _, ch := range word {
					// fmt.Println("hhh")
					res += asciiMap[ch][n]
				}
				res += "\n"
				// fmt.Println(res)
			}
		}
	}

	return res, nil
}
