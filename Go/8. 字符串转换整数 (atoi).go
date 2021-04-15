package Go

import "math"

func myAtoi(s string) int {
	var (
		index  int
		sign   bool  = false
		result int64 = 0
	)

	for index < len(s) && s[index] == ' ' {
		index++
	}
	if index < len(s) && s[index] == '+' {
		index++
	} else if index < len(s) && s[index] == '-' {
		index++
		sign = false
	}
	for index < len(s) {
		if s[index] <= '9' && s[index] >= '0' {
			result = result*10 + int64(s[index]-'0')
			if result > math.MaxInt32+1 {
				result = math.MaxInt32 + 1
			}
			index++
		} else {
			break
		}
	}
	if !sign {
		result = -result
	}
	if result > math.MaxInt32 {
		return math.MaxInt32
	} else if result < math.MinInt32 {
		return math.MinInt32
	} else {
		return int(result)
	}
}
