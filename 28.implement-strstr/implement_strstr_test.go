package algorithms28

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{",->0", args{"", ""}, 0},
		{"hello,ll->2", args{"hello", "ll"}, 2},
		{"aaaaa,bba->-1", args{"aaaaa", "bba"}, -1},
		{"aaaaa,aaa->0", args{"aaaaa", "aaa"}, 0},
		{"aaaaab,aab->3", args{"aaaaab", "aab"}, 3},
		{"aaa,aaaa->-1", args{"aaa", "aaaa"}, -1},
		{"aaaa,aaaa->0", args{"aaaa", "aaaa"}, 0},
		{"mississippi,issip->4", args{"mississippi", "issip"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
