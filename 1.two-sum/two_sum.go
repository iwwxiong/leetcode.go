package algorithms1

func twoSum(nums []int, target int) []int {
	var tempMap = make(map[int]int)
	var result []int
	for index, num := range nums {
		if i, ok := tempMap[target-num]; ok {
			result = append(result, i)
			result = append(result, index)
		} else {
			tempMap[num] = index
		}
	}

	return result
}
