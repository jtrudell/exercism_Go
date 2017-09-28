package pythagorean

import (
	"math"
)

const testVersion = 1

type Triplet [3]int

// Range returns a slice of Triplet with sides in the
// range min to max inclusive.
func Range(min, max int) []Triplet {
	var triplets []Triplet
	var squares []int

	for num := min; num <= max; num++ {
		squares = append(squares, num*num)
	}

	for i, valA := range squares {
		for j := i + 1; j < len(squares)-2; j++ {
			valB := squares[j]
			sum := valA + valB
			if findItem(squares[j+1:], sum) {
				c := math.Sqrt(float64(sum))
				triplets = append(triplets, Triplet{min + i, min + j, int(c)})
			}

		}
	}

	return triplets
}

// Sum returns a slice of Triplet where the sum a+b+c of each Triplet
// (the perimeter) is equal to p
func Sum(p int) []Triplet {
	var triplets []Triplet
	allTriplets := Range(1, p)
	for _, triplet := range allTriplets {
		if triplet[0]+triplet[1]+triplet[2] == p {
			triplets = append(triplets, triplet)
		}
	}
	return triplets
}

// findItem returns true if num is found in a slice of int and false otherwise
func findItem(x []int, num int) bool {
	for _, val := range x {
		if val == num {
			return true
		}
	}
	return false
}
