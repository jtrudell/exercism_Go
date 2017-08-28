package acronym

import "strings"

const testVersion = 3

func Abbreviate(phrase string) string {
	abbr := ""
	words := strings.Split(phrase, " ")

	for _, word := range words {
		letter := string(word[0])

		if strings.Contains(word, "-") {
			secondLetter := strings.Split(word, "-")[1][0]
			letter += string(secondLetter)
		}

		abbr += strings.ToUpper(letter)
	}
	return abbr
}
