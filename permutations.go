package main

import (
	"fmt"
)

type Candidates []int

func main() {
	var (
		candidates = []int{3, 4, 6}
	)

	c := permute(candidates)
	fmt.Println("c=", c)
}

func permute(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	var ret [][]int
	for i, n := range nums {
		sub := make([]int, len(nums)-1)
		copy(sub, nums[:i])
		copy(sub[i:], nums[i+1:])
		// 缩小到两个一组。
		for _, r := range permute(sub) {
			ret = append(ret, append([]int{n}, r...))
		}
	}
	return ret
}