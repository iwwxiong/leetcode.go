package algorithms85

import "testing"

func Test_maximalRectangle(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"0->1", args{[][]byte{{'0'}}}, 0},
		{"1->1", args{[][]byte{{'1'}}}, 1},
		{"11->1", args{[][]byte{{'1', '1'}}}, 2},
		{"0,0->1", args{[][]byte{{'0', '0'}}}, 0},
		{"10100,10111,11111,10010->6", args{[][]byte{
			{'1', '0', '1', '0', '0'},
			{'1', '0', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '0', '0', '1', '0'},
		}}, 6},
		{"", args{[][]byte{
			{'0', '0', '0', '0', '0', '0', '1'},
			{'0', '0', '0', '0', '1', '1', '1'},
			{'1', '1', '1', '1', '1', '1', '1'},
			{'0', '0', '0', '1', '1', '1', '1'},
		}}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalRectangle(tt.args.matrix); got != tt.want {
				t.Errorf("maximalRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
