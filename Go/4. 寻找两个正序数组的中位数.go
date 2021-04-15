package Go

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	M1 := ((len(nums1) + len(nums2) + 1) >> 1)
	M2 := ((len(nums1) + len(nums2) + 2) >> 1)
	if (len(nums1)+len(nums2))%2 == 0 {
		D1 := helper(nums1, 0, len(nums1)-1, nums2, 0, len(nums2)-1, M1)
		D2 := helper(nums1, 0, len(nums1)-1, nums2, 0, len(nums2)-1, M2)
		return (float64(D1) + float64(D2)) / 2
	} else {

		return float64(helper(nums1, 0, len(nums1)-1, nums2, 0, len(nums2)-1, M1))
	}

}

func helper(nums1 []int, lf1 int, rt1 int, nums2 []int, lf2 int, rt2 int, K int) int {
	len1 := rt1 - lf1 + 1
	len2 := rt2 - lf2 + 1
	if len1 > len2 {
		return helper(nums2, lf2, rt2, nums1, lf1, rt1, K)
	}
	if len1 == 0 {
		return nums2[lf2+K-1]
	}
	if K == 1 {
		return MinFun(nums1[lf1], nums2[lf2])
	}
	mid1 := lf1 + MinFun(K/2, len1) - 1
	mid2 := lf2 + MinFun(K/2, len2) - 1
	if nums1[mid1] < nums2[mid2] {
		return helper(nums1, mid1+1, rt1, nums2, lf2, rt2, K-(mid1-lf1+1))
	} else {
		return helper(nums1, lf1, rt1, nums2, mid2+1, rt2, K-(mid2-lf2+1))
	}
}

func MinFun(n1 int, n2 int) int {
	if n1 < n2 {
		return n1
	} else {
		return n2
	}
}
