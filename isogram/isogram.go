package isogram

import (
	"strings"
)

const testVersion = 1

func IsIsogram(s string) bool {
	letters := make(map[rune]int)
	for _, ltr := range strings.ToLower(s) {
		if letters[ltr] > 0 {
			return false
		}
		if ltr > 65 {
			letters[ltr]++
		}
	}
	return true
}
