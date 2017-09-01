package triangle

const testVersion = 3

type Kind string

var NaT Kind = "not a triangle"
var Equ Kind = "equilateral"
var Iso Kind = "isosceles"
var Sca Kind = "scalene"

func isTriangle(a, b, c float64) bool {
	return a+b > c && a+c > b && b+c > a
}

func KindFromSides(a, b, c float64) Kind {
	kind := Sca

	if !isTriangle(a, b, c) {
		kind = NaT
	} else if a == b && b == c {
		kind = Equ
	} else if a == b || b == c || a == c {
		kind = Iso
	}

	return kind
}