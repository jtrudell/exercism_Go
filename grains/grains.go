package grains

import "errors"

const testVersion = 1

func Square(n int) (uint64, error) {
	var err error
	squares := make([]uint64, 64)

	if n < 1 || n > 64 {
		err = errors.New("Input must be an integer between 1 and 64.")
		return squares[0], err
	}

	squares[0] = uint64(1)
	for i := 1; i < n; i++ {
		squares[i] = squares[i-1] * 2
	}

	return squares[n-1], err
}

func Total() uint64 {
	var sum uint64

	for i := 1; i <= 64; i++ {
		value, _ := Square(i)
		sum += value
	}

	return sum
}
