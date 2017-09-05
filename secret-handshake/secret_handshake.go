package secret

const testVersion = 2

func Reverse(s []string) []string {
	j := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
	return s
}

func Handshake(num uint) []string {
	res := []string{}
	reverse := true

	if num >= 16 {
		reverse = false
		num %= 16
	}

	if num >= 8 {
		res = append(res, "jump")
		num %= 8
	}

	if num >= 4 {
		res = append(res, "close your eyes")
		num %= 4
	}

	if num >= 2 {
		res = append(res, "double blink")
		num %= 2
	}

	if num == 1 {
		res = append(res, "wink")
	}

	if reverse && len(res) > 1 {
		res = Reverse(res)
	}

	return res
}
