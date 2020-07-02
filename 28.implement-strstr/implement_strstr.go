package algorithms28

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}

	var j int = -1
	for i := 0; i <= len(haystack); i++ {
		if len(haystack[i:]) >= len(needle) {
			if haystack[i:i+len(needle)] == needle {
				j = i
				break
			}
		}
	}
	return j
}
