package algorithms3

func lengthOfLongestSubstring(s string) int {
	var start = -1
	var tempMap = make(map[string]int)
	var maxLen = 0
	for i, chr := range s {
		index, ok := tempMap[string(chr)]
		if ok {
			if start < index {
				start = index
			}
		}
		if i-start > maxLen {
			maxLen = i - start
		}
		tempMap[string(chr)] = i
	}
	return maxLen
}
