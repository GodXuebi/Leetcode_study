package Go

import "strings"

func convert(s string, numRows int) string {
	if numRows == 0 {
		return ""
	}
	if numRows == 1 {
		return s
	}
	if len(s) <= numRows {
		return s
	}
	Row := make([]string, numRows)
	goingdown := false
	cur_row := 0

	for i := 0; i < len(s); i++ {
		Row[cur_row] += string(s[i])
		if cur_row == 0 || cur_row == numRows-1 {
			goingdown = !goingdown
		}
		if goingdown == true {
			cur_row++
		} else {
			cur_row--
		}
	}
	return strings.Join(Row, "")
}
