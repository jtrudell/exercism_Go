package summultiples

const testVersion = 2

func SumMultiples(limit int, divisors ...int) int {
	var sum int
	multiples := make(map[int]int)

	for i, val := range divisors {
		for val < limit {
			if _, ok := multiples[val]; !ok {
				multiples[val] = sum + val
				sum = multiples[val]
			}
			val += divisors[i]
		}
	}

	return sum
}
