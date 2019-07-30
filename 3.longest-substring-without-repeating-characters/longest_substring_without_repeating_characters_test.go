package algorithms3

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	// TODO: Add test cases.
		{"a->1", args{"a"}, 1},
		{"bbbbb->1", args{"bbbbb"}, 1},
		{"abcabcbb->3", args{"abcabcbb"}, 3},
		{"abcdac->4", args{"abcdac"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
