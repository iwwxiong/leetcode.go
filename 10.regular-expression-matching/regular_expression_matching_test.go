package algorithms10

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"aa,a->false", args{"aa", "a"}, false},
		{"aa,aa->true", args{"aa", "aa"}, true},
		{"aaa,aa->false", args{"aaa", "aa"}, false},
		{"aa,a*->true", args{"aa", "a*"}, true},
		{"aa,.*->true", args{"aa", ".*"}, true},
		{"ab,.*->true", args{"ab", ".*"}, true},
		{"aab,c*a*b->true", args{"aab", "c*a*b"}, true},
		{"mississippi,mis*is*p", args{"mississippi", "mis*is*p"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
