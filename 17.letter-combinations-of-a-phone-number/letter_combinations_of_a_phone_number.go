package algorithms17

func merge(as, bs []string) []string {
	var data []string
	for _, a := range as {
		for _, b := range bs {
			data = append(data, a+b)
		}
	}
	return data
}

func letterCombinations(digits string) []string {
	var letterMap map[string][]string = map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
	var ls [][]string
	for _, digit := range digits {
		letter := string(digit)
		ls = append(ls, letterMap[letter])
	}
	if len(ls) == 0 {
		return []string{}
	} else if len(ls) == 1 {
		return ls[0]
	} else {
		var data []string = ls[0]
		for _, l := range ls[1:] {
			data = merge(data, l)
		}
		return data
	}
}
