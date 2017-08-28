package accumulate

const testVersion = 1

func Accumulate(x []string, y func(string) string) []string {
	for i, value := range x {
		x[i] = y(value)
	}
	return x
}
