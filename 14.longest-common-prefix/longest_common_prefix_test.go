package algorithms14

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"[aa,a]->a", args{[]string{"aa", "a"}}, "a"},
		{"[abc,abc,abc]->abc", args{[]string{"abc", "abc", "abc"}}, "abc"},
		{"[flower,flow,flight]->fl", args{[]string{"flower", "flow", "flight"}}, "fl"},
		{"[dog,racecar,car]->fl", args{[]string{"dog", "racecar", "car"}}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
