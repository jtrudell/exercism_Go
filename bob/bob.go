package bob

import (
	"regexp"
	"strings"
)

const testVersion = 3

func Hey(s string) string {
	nonAlpha := regexp.MustCompile(`[\d,;<>./'"{}()-+=%#&@*^*$\s]`)
	allCaps := regexp.MustCompile(`^([A-Z][^a-z]\s*[!,?]*)+$`)

	alphaStr := nonAlpha.ReplaceAllString(s, "")

	res := "Whatever."

	if strings.TrimSpace(s) == "" {
		res = "Fine. Be that way!"
	} else if allCaps.MatchString(alphaStr) {
		res = "Whoa, chill out!"
	} else if strings.HasSuffix(alphaStr, "?") {
		res = "Sure."
	}

	return res
}
