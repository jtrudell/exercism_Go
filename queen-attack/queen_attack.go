package queenattack

import (
	"errors"
	"strconv"
	"strings"
)

const testVersion = 2

func Stoi(s []string) (int, int, error) {
	var err error
	var n int

	l, convErr := strconv.Atoi(s[1])
	if convErr != nil {
		err = convErr
	}

	switch s[0] {
	case "a":
		n = 1
	case "b":
		n = 2
	case "c":
		n = 3
	case "d":
		n = 4
	case "e":
		n = 5
	case "f":
		n = 6
	case "g":
		n = 7
	case "h":
		n = 8
	default:
		err = errors.New("out of range")
	}
	return n, l, err
}

func CanQueenAttack(w, b string) (attack bool, err error) {
	attack = false

	// return with error if w and b arguments are invalid
	if w == b || len(w) != 2 || len(b) != 2 {
		err = errors.New("same spaces given or no space given")
		return
	}

	// convert letter and string int to ints
	wLtr, wNum, wErr := Stoi(strings.Split(w, ""))
	bLtr, bNum, bErr := Stoi(strings.Split(b, ""))

	// return with error if arguments are out of range
	if wErr != nil || wNum < 1 || wNum > 8 || wLtr > 8 {
		err = errors.New("out of range")
		return
	}

	if bErr != nil || bNum < 1 || bNum > 8 || bLtr > 8 {
		err = errors.New("out of range")
		return
	}

	// check same row or column or swap row/column
	if (wLtr == bLtr || wNum == bNum) || (wLtr == bNum && wNum == bLtr) {
		attack = true
		return
	}

	// check diagonals left
	for i := 1; wNum-i > 0 && wLtr-i > 0; i++ {
		if wNum-i == bNum && wLtr-i == bLtr {
			attack = true
			return
		}
	}

	// check diagonals right
	for i := 1; wNum+i < 9 && wLtr+i < 9; i++ {
		if wNum+i == bNum && wLtr+i == bLtr {
			attack = true
			return
		}
	}

	return
}
