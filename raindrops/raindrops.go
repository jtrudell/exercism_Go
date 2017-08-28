package raindrops

import "strconv"

const testVersion = 3

func Convert(num int) string {
	rain := ""

	if num%3 == 0 {
		rain += "Pling"
	}

	if num%5 == 0 {
		rain += "Plang"
	}

	if num%7 == 0 {
		rain += "Plong"
	}

	if rain == "" {
		rain = strconv.Itoa(num)
	}

	return rain
}
