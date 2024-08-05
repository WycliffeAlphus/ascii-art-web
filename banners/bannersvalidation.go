package banners

import (
	"crypto/sha1"
	"reflect"
)

// IsValidBanner ensures that banner files are not altered
func IsValidBanner(patternFile string, textFile []byte) bool {
	if patternFile == "banners/shadow.txt" && !reflect.DeepEqual(sha1.Sum(textFile), [20]byte{88, 40, 71, 6, 150, 77, 77, 228, 230, 34, 145, 25, 177, 146, 82, 213, 194, 159, 105, 87}) {
		return false
	}

	if patternFile == "banners/standard.txt" && sha1.Sum(textFile) != [20]byte{56, 168, 206, 246, 119, 46, 15, 119, 221, 52, 31, 98, 190, 17, 74, 113, 102, 243, 28, 92} {
		return false
	}
	if patternFile == "banners/thinkertoy" && sha1.Sum(textFile) != [20]byte{102, 110, 218, 73, 92, 48, 212, 164, 41, 25, 65, 110, 169, 170, 207, 96, 248, 42, 244, 253} {
		return false
	}
	return true
}
