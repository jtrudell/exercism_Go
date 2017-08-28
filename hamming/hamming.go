package hamming

import (
	"errors"
)

const testVersion = 6

func Distance(a, b string) (int, error) {
	var err error
	count := 0

	if len(a) != len(b) {
		err = errors.New("String lengths do not match")
		return count, err
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count, err
}
