package algorithms14

func longestCommonPrefix(strs []string) string {
	var prefix string
	if len(strs) == 0 {
		return prefix
	}
	var first string = strs[0]
	var flag bool = true
	var i, length int = 0, len(first)
	for flag && i < length {
		for _, str := range strs[1:] {
			if i >= len(str) || first[i] != str[i] {
				flag = false
				break
			}
		}
		if flag {
			i++
		}
	}
	return first[:i]
}
