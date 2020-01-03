package algorithms13

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"III->3", args{"III"}, 3},
		{"IV->4", args{"IV"}, 4},
		{"IX->9", args{"IX"}, 9},
		{"LVIII->58", args{"LVIII"}, 58},
		{"MCMXCIV->1994", args{"MCMXCIV"}, 1994},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
