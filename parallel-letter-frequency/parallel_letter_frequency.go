package letter

const testVersion = 1

// In which I write my own frequency() function because I didn't read the directions
func ConcurrentFrequency(s []string) map[rune]int {
	freq := make(map[rune]int)
	counts := make([]map[rune]int, len(s))

	for i, str := range s {
		counts[i] = frequency(str)
	}

	for _, count := range counts {
		for k, v := range count {
			freq[k] += v
		}
	}

	return freq
}

func frequency(str string) map[rune]int {
	freq := make(map[rune]int)
	letters := []rune(str)

	for _, letter := range letters {
		freq[letter]++
	}

	return freq
}
