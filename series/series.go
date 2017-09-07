package series

const testVersion = 2

func All(n int, s string) []string {
	var strSlice []string
	strEnd := len(s)

	for i := 0; i <= strEnd-n; i++ {
		strSlice = append(strSlice, UnsafeFirst(n, string(s[i:strEnd])))
	}
	return strSlice
}

func UnsafeFirst(n int, s string) string {
	return string(s[0:n])
}
