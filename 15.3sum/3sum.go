package algorithms15

func threeSum(nums []int) [][]int {
	var length int = len(nums)
	var result [][]int
	if length < 3 {
		return result
	}

	tm := make(map[int][]int)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length-1; j++ {
			num := 0 - (nums[i] + nums[j])
			if _, ok := tm[num]; !ok {
				tm[num] = []int{nums[i], nums[j]}
			}
		}
	}

	for _, num := range nums[2:] {
		if s, ok := tm[num]; ok {
			s := append(s, num)
			result = append(result, s)
			delete(tm, num)
		}
	}

	return result
}
