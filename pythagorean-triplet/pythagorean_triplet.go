package pythagorean

// The three elements of each returned triplet must be in order,
// t[0] <= t[1] <= t[2], and the list of triplets must be in lexicographic
// order.

const testVersion = 1

type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) []Triplet {
	x := []Triplet{Triplet{1, 2, 3}}

	return x
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(p int) []Triplet {
	x := []Triplet{Triplet{1, 2, 3}}
	return x
}
