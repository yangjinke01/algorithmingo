package sort

import (
	"algorithmingo/utils"
	"testing"
)

type args struct {
	nums []int
}
type want struct {
	nums []int
}

var tests = []struct {
	name string
	args args
	want want
}{
	{name: "normal", args: args{nums: []int{3, 4, 1}}, want: want{nums: []int{1, 3, 4}}},
	{name: "normal", args: args{nums: []int{4, 2, 1}}, want: want{nums: []int{1, 2, 4}}},
	{name: "negative", args: args{nums: []int{-3, 4, -1, 0}}, want: want{nums: []int{-3, -1, 0, 4}}},
}

func Test_Sort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//selectSort(tt.args.nums)
			//bubbleSort(tt.args.nums)
			insertSort(tt.args.nums)
			if !utils.Equal(tt.args.nums, tt.want.nums) {
				t.Errorf("got %v, want %v", tt.args.nums, tt.want.nums)
			}
		})
	}
}
