package algorithms4

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var j int
	var median float64
	var length = len(nums1) + len(nums2)
	var tempNum []int
	for _, num1 := range nums1 {
		if len(tempNum) >= length/2+1 {
			break
		}
		for j < len(nums2) && nums2[j] <= num1 {
			tempNum = append(tempNum, nums2[j])
			j++
		}
		tempNum = append(tempNum, num1)
	}
	if len(nums2[j:]) > 0 && len(tempNum) < length/2+1 {
		tempNum = append(tempNum, nums2[j:]...)
	}
	if length%2 == 1 {
		median = float64(tempNum[length/2])
	} else {
		median = float64(tempNum[length/2-1]+tempNum[length/2]) / float64(2)
	}

	return median
}
