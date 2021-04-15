package Go

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	map1 := make(map[byte]int, 26)
	left, right := 0, 0
	minlen := 1

	for right < len(s) {
		ch1 := s[right]
		map1[ch1]++
		right++
		for map1[ch1] > 1 {
			ch2 := s[left]
			map1[ch2]--
			left++
		}
		if right-left > minlen {
			minlen = right - left
		}
	}
	return minlen
}
