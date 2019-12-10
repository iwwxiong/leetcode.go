package althorism8

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"42->42", args{"42"}, 42},
		{"4a2b->42", args{"4a2b"}, 42},
		{"-42->-42", args{"-42"}, -42},
		{"+-2->0", args{"+-2"}, 0},
		{"3.14159->3", args{"3.14159"}, 3},
		{"     -42->-42", args{"     -42"}, -42},
		{"  \r  -42->-42", args{"  \r  -42"}, -42},
		{"4193 with words->4193", args{"4193 with words"}, 4193},
		{"words and 987->0", args{"words and 987"}, 0},
		{"-91283472332->-2147483648", args{"-91283472332"}, -2147483648},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
