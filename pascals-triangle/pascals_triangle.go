package pascal

const testVersion = 1

func Triangle(n int) [][]int {
	var pascal = [][]int{{1}}

	for i := 1; i < n; i++ {
		row := make([]int, i+1)
		last := len(row) - 1
		row[0], row[last] = 1, 1

		for j := 1; j < last; j++ {
			row[j] = pascal[i-1][j-1] + pascal[i-1][j]
		}
		pascal = append(pascal, row)
	}
	return pascal
}
