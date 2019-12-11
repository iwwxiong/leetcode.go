package algorithms8

import "math"

func myAtoi(str string) int {
	var sign string
	var i, total int

	for _, ch := range str {
		if total > math.MaxInt32 {
			break
		}

		if i == 0 {
			if string(ch) == " " || string(ch) == "\t" || string(ch) == "\n" || string(ch) == "\f" || string(ch) == "\b" || string(ch) == "\r" {
				continue
			}

			if string(ch) == "-" || string(ch) == "+" {
				if sign != "" {
					return 0
				}
				i++
				sign = string(ch)
				continue
			}
		}

		if string(ch) >= "0" && string(ch) <= "9" {
			i++
			total = 10*total + int(ch) - 48
		} else {
			break
		}

	}

	if total > math.MaxInt32 {
		if sign == "-" {
			return math.MinInt32
		}
		return math.MaxInt32
	}

	if sign == "-" {
		return -int(total)
	}
	return int(total)
}
