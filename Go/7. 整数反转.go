package Go

import "math"

func reverse(x int) int {
	y := 0
	for x != 0 {
		y = y*10 + x%10
		x = x / 10
	}
	if y > math.MaxInt32 || y < math.MinInt32 {
		return 0
	} else {
		return y
	}
}
