package algorithms12

func intToRoman(num int) string {
	var romanString string
	if num < 0 || num > 3999 {
		panic("num range from 0, 3999")
	}

	appendS := func(roman string, index int, flags ...string) string {
		if len(flags) == 2 {
			roman += (flags[0] + flags[1])
		} else {
			for i := 0; i < index; i++ {
				roman += flags[0]
			}
		}
		return roman
	}

	if i := num / 1000; i > 0 {
		num = num % 1000
		romanString = appendS(romanString, i, "M")
	}

	if i := num / 900; i > 0 {
		num = num % 900
		romanString = appendS(romanString, i, "C", "M")
	}

	if i := num / 500; i > 0 {
		num = num % 500
		romanString = appendS(romanString, i, "D")
	}

	if i := num / 400; i > 0 {
		num = num % 400
		romanString = appendS(romanString, i, "C", "D")
	}

	if i := num / 100; i > 0 {
		num = num % 100
		romanString = appendS(romanString, i, "C")
	}

	if i := num / 90; i > 0 {
		num = num % 90
		romanString = appendS(romanString, i, "X", "C")
	}

	if i := num / 50; i > 0 {
		num = num % 50
		romanString = appendS(romanString, i, "L")
	}

	if i := num / 40; i > 0 {
		num = num % 40
		romanString = appendS(romanString, i, "X", "L")
	}

	if i := num / 10; i > 0 {
		num = num % 10
		romanString = appendS(romanString, i, "X")
	}

	if i := num / 9; i > 0 {
		num = num % 9
		romanString = appendS(romanString, i, "I", "X")
	}

	if i := num / 5; i > 0 {
		num = num % 5
		romanString = appendS(romanString, i, "V")
	}

	if num > 0 {
		if num == 4 {
			romanString = appendS(romanString, num, "I", "V")
		} else {
			romanString = appendS(romanString, num, "I")
		}
	}

	return romanString
}
