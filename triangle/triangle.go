package triangle

import "math"

const testVersion = 3

type Kind string

var NaT Kind = "not a triangle"
var Equ Kind = "equilateral"
var Iso Kind = "isosceles"
var Sca Kind = "scalene"

func isTriangle(a, b, c float64) bool {
	return a+b >= c && a+c >= b && b+c >= a
}

func infSide(a, b, c float64) bool {
	return math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0)
}

func zeroSide(a, b, c float64) bool {
	return a <= 0 || b <= 0 || c <= 0
}

func KindFromSides(a, b, c float64) Kind {
	if !isTriangle(a, b, c) || infSide(a, b, c) || zeroSide(a, b, c) {
		return NaT
	} else if a == b && b == c {
		return Equ
	} else if a == b || b == c || a == c {
		return Iso
	}

	return Sca
}
