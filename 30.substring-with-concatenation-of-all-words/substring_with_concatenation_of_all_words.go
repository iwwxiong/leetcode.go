package algorithms30

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}
	var wordLength int = len(words[0])
	if len(s) < wordLength {
		return []int{}
	}
	var wordMap map[string]int = make(map[string]int)
	for _, word := range words {
		_, ok := wordMap[word]
		if ok {
			wordMap[word] = wordMap[word] + 1
		} else {
			wordMap[word] = 1
		}
	}
	var result []int = []int{}
	for i := 0; i < len(s)-len(words)*wordLength+1; i++ {
		var wordMapV2 map[string]int = make(map[string]int)
		var j int
		for j = 0; j < len(words); j++ {
			var word string = s[i+j*wordLength : i+(j+1)*wordLength]
			num, ok := wordMap[word]
			if !ok {
				break
			}
			numV2, ok := wordMapV2[word]
			if ok {
				numV2++
			} else {
				numV2 = 1
			}
			wordMapV2[word] = numV2
			if numV2 > num {
				break
			}
		}
		if j == len(words) {
			result = append(result, i)
		}
	}
	return result
}
