package letter

import "sync"

const testVersion = 1

var wg sync.WaitGroup

func ConcurrentFrequency(s []string) FreqMap {
	freq := FreqMap{}
	counts := make([]FreqMap, len(s))
	wg.Add(len(s))

	for i, str := range s {
		go func(str string, i int) {
			counts[i] = Frequency(str)
			wg.Done()
		}(str, i)
	}
	wg.Wait()

	for _, count := range counts {
		for k, v := range count {
			freq[k] += v
		}
	}

	return freq
}
