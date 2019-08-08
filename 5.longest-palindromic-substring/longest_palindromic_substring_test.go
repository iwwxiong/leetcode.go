package algorithms5

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"a->a", args{"a"}, "a"},
		{"aa->a", args{"ac"}, "a"},
		{"bb->bb", args{"bb"}, "bb"},
		{"ccd->cc", args{"ccd"}, "cc"},
		{"cbbd->bb", args{"cbbd"}, "bb"},
		{"abcba->abcba", args{"abcba"}, "abcba"},
		{"abbacc->abba", args{"abbacc"}, "abba"},
		{"aabcacb->bcacb", args{"aabcacb"}, "bcacb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
