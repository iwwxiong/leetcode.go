package algorithms38

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	var prev string = countAndSay(n - 1)
	return nextString(prev)
}

func nextString(prev string) string {
	if len(prev) == 0 {
		return ""
	}
	var first string = string(prev[0])
	num := repeatNum(prev, first)
	return simpleIntToString(num) + first + nextString(prev[num:len(prev)])
}

func simpleIntToString(n int) string {
	var s string
	switch n {
	case 1:
		s = "1"
	case 2:
		s = "2"
	case 3:
		s = "3"
	case 4:
		s = "4"
	case 5:
		s = "5"
	}
	return s
}

func repeatNum(char string, s string) int {
	var num int = 0
	for _, i := range char {
		if string(i) == s {
			num++
		} else {
			break
		}
	}
	return num
}
