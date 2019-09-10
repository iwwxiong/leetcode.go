package algorithms6

func convert(s string, numRows int) string {
	if len(s) <= 1 || numRows == 1 || numRows >= len(s) {
		return s
	}

	type intSlice []string
	var result string
	var tempS = make([]intSlice, numRows)
	var interval = 2*numRows - 2
	var step int
	if len(s)%interval == 0 {
		step = len(s) / interval
	} else {
		step = len(s)/interval + 1
	}
	for i := 0; i < step; i++ {
		var start, end = i * interval, (i + 1) * interval
		if (i+1)*interval > len(s) {
			end = len(s)
		}
		for j, s := range s[start:end] {
			if j < numRows {
				tempS[j] = append(tempS[j], string(s))
			} else {
				tempS[2*(numRows-1)-j] = append(tempS[2*(numRows-1)-j], string(s))
			}
		}
	}

	for _, ts := range tempS {
		for _, str := range ts {
			result += str
		}
	}
	return result
}
