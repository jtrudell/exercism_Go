package pangram

import "strings"

const testVersion = 2

func IsPangram(str string) bool {
	if len(str) < 26 {
		return false
	}

	str = strings.ToLower(str)

	for i := int('a'); i <= int('z'); i++ {
		if !strings.ContainsRune(str, rune(i)) {
			return false
		}
	}

	return true
}
