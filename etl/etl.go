package etl

import "strings"

const testVersion = 1

func Transform(input map[int][]string) map[string]int {
	output := make(map[string]int)

	for key, value := range input {
		for _, str := range value {
			output[strings.ToLower(str)] = key
		}
	}
	return output
}
