package Go

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return s
	}

	maxleft, maxright := 0, 0
	maxlen := 1
	dp := make([][]bool, len(s))

	for right := 0; right < len(s); right++ {
		dp[right] = make([]bool, len(s))
		for left := 0; left <= right; left++ {
			if (right-left+1 <= 2 || dp[left+1][right-1] == true) && s[left] == s[right] {
				dp[left][right] = true
				if right-left+1 >= maxlen {
					maxleft = left
					maxright = right
					maxlen = right - left + 1
				}
			} else {
				dp[left][right] = false
			}
		}
	}
	return s[maxleft : maxright+1]
}
