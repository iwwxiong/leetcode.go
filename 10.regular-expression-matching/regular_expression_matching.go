package algorithms10

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	if len(p) == 1 || string(p[1]) != "*" {
		if len(s) == 0 {
			return false
		}

		if string(p[0]) != "." && string(p[0]) != string(s[0]) {
			return false
		}

		return isMatch(s[1:], p[1:])
	}

	var i int = -1
	for i < len(s) && (i < 0 || string(p[0]) == "." || s[i] == p[0]) {
		if isMatch(s[i+1:], p[2:]) {
			return true
		}
		i++
	}

	return false
}
