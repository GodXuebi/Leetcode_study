package Go

func twoSum(nums []int, target int) []int {
	mymap := map[int]int{}
	for i, k := range nums {
		if v, ok := mymap[target-k]; ok {
			return []int{i, v}
		}
		mymap[k] = i
	}
	return []int{}
}
