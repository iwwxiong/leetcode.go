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
		{"4a2b->42", args{"4a2b"}, 4},
		{"-42->-42", args{"-42"}, -42},
		{"+-2->0", args{"+-2"}, 0},
		{"3.14159->3", args{"3.14159"}, 3},
		{"     -42->-42", args{"     -42"}, -42},
		{"  \r  -42->-42", args{"  \r  -42"}, -42},
		{"   +0 123", args{"   +0 123"}, 0},
		{"  -0012a42", args{"  -0012a42"}, -12},
		{"4193 with words->4193", args{"4193 with words"}, 4193},
		{"words and 987->0", args{"words and 987"}, 0},
		{"-   234", args{"-   234"}, 0},
		{"2147483648", args{"2147483648"}, 2147483647},
		{"-91283472332->-2147483648", args{"-91283472332"}, -2147483648},
		{"9223372036854775808", args{"9223372036854775808"}, 2147483647},
		{"18446744073709551617", args{"18446744073709551617"}, 2147483647},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
