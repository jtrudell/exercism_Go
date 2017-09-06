package series

import "strings"

const testVersion = 2

func All(n int, s string) []string {
	var strSlice []string
	strEnd := len(s)

	for i := 0; i <= strEnd-n; i++ {
		substr := strings.Split(s, "")
		next := UnsafeFirst(n, strings.Join(substr[i:strEnd], ""))
		strSlice = append(strSlice, next)
	}
	return strSlice
}

func UnsafeFirst(n int, s string) string {
	substr := strings.Split(s, "")
	return strings.Join(substr[0:n], "")
}
