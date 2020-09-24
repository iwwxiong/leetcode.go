package algorithms88

func mergeV1(nums1 []int, m int, nums2 []int, n int) []int {
	var i, j int = m - 1, n - 1
	var num []int = make([]int, m+n)
	for i >= 0 && j >= 0 {
		if nums1[i] <= nums2[j] {
			num[i+j+1] = nums2[j]
			j--
		} else {
			num[i+j+1] = nums1[i]
			i--
		}
	}
	for i >= 0 {
		num[i] = nums1[i]
		i--
	}
	for j >= 0 {
		num[j] = nums2[j]
		j--
	}
	nums1 = num
	return nums1
}

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	var i, j int = m - 1, n - 1
	var k int = i + j + 1
	for k >= 0 {
		if j < 0 {
			break
		}
		if i < 0 {
			nums1[k] = nums2[j]
			j--
		} else if nums1[i] < nums2[j] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums1[i]
			i--
		}
		k--
	}
	return nums1
}
