package artgen

import (
	"errors"
	"os"
	"strings"
)

// AsciiMapping given a banner file, reads all graphics representations of the ASCII characters and
// returns a map of the ASCII character to the graphics representations of the ASCII character
func AsciiMapping(patternFile string) (map[rune][]string, error) {
	var splitted []string

	switch patternFile {
	case "standard":
		patternFile = "banners/standard.txt"
	case "shadow":
		patternFile = "banners/shadow.txt"
	case "thinkertoy":
		patternFile = "banners/thinkertoy.txt"
	}

	textFile, err := os.ReadFile(patternFile)
	if err != nil {
		return nil, err
	}

	if len(textFile) == 0 {
		err = errors.New("provided banner file is empty")
		return nil, err
	}

	switch patternFile {
	case "banners/thinkertoy.txt":
		splitted = strings.Split(string(textFile), "\r\n") // strings of thinkeratoi are seperated by \r\n [13,10]
	default:
		splitted = strings.Split(string(textFile), "\n")

	}
	if len(splitted) != 856 {
		err = errors.New("provided banner file is corrupted")
		return nil, err
	}

	asciiMap := make(map[rune][]string)
	startAscii := ' '
	for i := 1; i < len(splitted); {
		arrayString := []string{}
		for j := 0; j < 8; j++ {
			arrayString = append(arrayString, splitted[i])
			i++
		}
		i++
		asciiMap[startAscii] = arrayString
		startAscii++
	}

	return asciiMap, nil
}
