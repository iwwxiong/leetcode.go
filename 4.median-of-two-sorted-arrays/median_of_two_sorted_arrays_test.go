package algorithms4

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"[1,2],[3]->2", args{[]int{1, 2}, []int{3}}, 2},
		{"[1,2],[3,4]->2.5", args{[]int{1, 2}, []int{3, 4}}, 2.5},
		{"[1,2,3],[3,4]->2.5", args{[]int{1, 2, 3}, []int{3, 4}}, 3},
		{"[4,5],[1,2,3]->3", args{[]int{4, 5}, []int{1, 2, 3}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
