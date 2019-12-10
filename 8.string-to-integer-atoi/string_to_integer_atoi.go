package althorism8

func myAtoi(str string) int {
	var sign bool = true
	var i, total int

	for _, ch := range str {
		if string(ch) == " " || string(ch) == "\t" || string(ch) == "\n" || string(ch) == "\f" || string(ch) == "\b" || string(ch) == "\r" {
			continue
		}

		if i == 0 {
			if string(ch) == "-" {
				sign = false
				continue
			}

			if string(ch) == "+" {
				continue
			}

			if string(ch) < "0" || string(ch) > "9" {
				return 0
			}
		}

		if string(ch) == "." {
			break
		}

		if string(ch) >= "0" && string(ch) <= "9" {
			total = 10*total + (int(ch) - 48)
		}

	}

	if !sign {
		total = -total
	}

	if total > 2147483647 {
		total = 2147483648
	} else if total < -2147483648 {
		total = -2147483648
	}

	return total
}
