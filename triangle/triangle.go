// Package triangle exposes KindFromSides function.
package triangle

import "math"

type Kind int

const (
    NaT = iota // not a triangle
    Equ // equilateral
    Iso // isosceles
    Sca // scalene
)

// KindFromSides returns triangle Kind given 3 of its sides.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	ch := make (chan bool, 1)
	go assertSumLarger(a, b, c, ch)
	if a <= 0 || b <= 0 || c <= 0 || math.IsNaN(a+b+c) || math.IsInf(a+b+c, 1) {
		return NaT
	} else if a == b && b == c {
		k = Equ
	} else if a == b || a == c || b == c {
		k = Iso
	} else {
		k = Sca
	}
	larger := <- ch
	if !larger {
		k = NaT
	}
	return k
}

func assertSumLarger(a, b, c float64, ch chan bool) {
	larger := true
	sum := a + b + c
	sides := []float64{a, b, c}
	for _, s := range sides {
		if sum - s < s {
			larger = false
		}
	}
	ch <- larger
	return
}
