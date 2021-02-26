package pythagorean

type Triplet struct {
	a, b, c int
}

func Range(min, max int) []Triplet {
	triplets := []Triplet{}
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			for c := b; c <= max; c++ {
				if a*a+b*b == c*c {
					triplets = append(triplets, Triplet{a, b, c})
				}
			}
		}
	}
	return triplets
}

func Sum(sum int) []Triplet {
	triplets := []Triplet{}
	for a := 1; a <= sum/2; a++ {
		for b := a; b <= sum/2; b++ {
			c := sum - a - b
			if b > c {
				break
			}
			if a*a+b*b == c*c {
				triplets = append(triplets, Triplet{a, b, c})
			}
		}
	}
	return triplets
}
