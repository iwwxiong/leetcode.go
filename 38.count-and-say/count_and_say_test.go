package algorithms38

import "testing"

func Test_countAndSay(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1->1", args{1}, "1"},
		{"2->11", args{2}, "11"},
		{"3->21", args{3}, "21"},
		{"4->1211", args{4}, "1211"},
		{"5->111221", args{5}, "111221"},
		{"6->312211", args{6}, "312211"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAndSay(tt.args.n); got != tt.want {
				t.Errorf("countAndSay() = %v, want %v", got, tt.want)
			}
		})
	}
}
