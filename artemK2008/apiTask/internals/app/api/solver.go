package api

import (
	"math"
	"strconv"
)

func countRoots(a1, b1, c1 string) int {
	a, _ := strconv.Atoi(a1)
	b, _ := strconv.Atoi(b1)
	c, _ := strconv.Atoi(c1)
	if a == 0 && b == 0 && c != 0 {
		return 0
	}

	if a == 0 || (b == 0 && c == 0) {
		return 1
	}

	if b == 0 {
		if c/a < 0 {
			return 0
		}
	}
	var discr int = b*b - 4*a*c
	if discr < 0 {
		return 0
	}
	if discr == 0 {
		return 1
	}
	i := int(math.Sqrt(float64(discr)))
	if (b*(-1)+i)/2*a == (b*(-1)-i)/2*a {
		return 1
	}
	return 2
}
