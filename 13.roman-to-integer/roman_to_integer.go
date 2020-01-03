package algorithms13

func romanToInt(s string) int {
	var i, romanInt int
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	var intSlice []int
	for _, roman := range s {
		intSlice = append(intSlice, romanMap[string(roman)])
	}
	length := len(intSlice)
	for i < length {
		if i == length-1 {
			romanInt += intSlice[i]
			i++
		} else if intSlice[i+1] > intSlice[i] {
			romanInt += (intSlice[i+1] - intSlice[i])
			i += 2
		} else {
			romanInt += intSlice[i]
			i++
		}
	}
	return romanInt
}
