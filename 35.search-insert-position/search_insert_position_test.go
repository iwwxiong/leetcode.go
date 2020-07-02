package algorithms35

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1,3,5,6 5->2", args{[]int{1, 3, 5, 6}, 5}, 2},
		{"1,3,5,6 2->1", args{[]int{1, 3, 5, 6}, 2}, 1},
		{"1,3,5,6 7->4", args{[]int{1, 3, 5, 6}, 7}, 4},
		{"1,3,5,6 5->0", args{[]int{1, 3, 5, 6}, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
