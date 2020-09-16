package algorithms67

func stoi(s string) int {
	if s == "1" {
		return 1
	} else {
		return 0
	}
}

func itos(i int) string {
	if i == 1 {
		return "1"
	} else {
		return "0"
	}
}

func addBinary(a string, b string) string {
	var i int = len(a) - 1
	var j int = len(b) - 1
	var s string = ""
	var carry int = 0
	for i >= 0 || j >= 0 {
		var n1, n2 int = 0, 0
		if i >= 0 {
			n1 = stoi(string(a[i]))
			i--
		}
		if j >= 0 {
			n2 = stoi(string(b[j]))
			j--
		}
		c := n1 + n2 + carry
		if c >= 2 {
			carry = 1
			s = itos(c%2) + s
		} else {
			s = itos(c) + s
			carry = 0
		}
	}
	if carry == 1 {
		s = itos(carry) + s
	}
	return s
}
