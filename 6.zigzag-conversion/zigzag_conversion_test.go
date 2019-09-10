package algorithms6

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"0123456789,4->0615724839", args{"0123456789", 4}, "0615724839"},
		{"PAYPALISHIRING,3->PAHNAPLSIIGYIR", args{"PAYPALISHIRING", 3}, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING,4->PINALSIGYAHRPI", args{"PAYPALISHIRING", 4}, "PINALSIGYAHRPI"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
