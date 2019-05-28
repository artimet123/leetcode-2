package main

import (
	"fmt"
)

type Candidates []int

func main() {
	var (
		candidates = []int{3, 4, 6, 2, 7, 8, 10, 13}
		target     int
	)

	target = firstMissingPositive(candidates)
	fmt.Println("target=", target)
}

func firstMissingPositive(nums []int) int {
	n := len(nums) // n = 9

	for i := 0; i < n; i++ {
		if j := nums[i] - 1; j >= 0 && j < n {
			if nums[i] != nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
				i-- // 让i一直保持在该位置上，比较新的值和新的别的值的比较。
			}
		}
	}

	for i := 0; i < n; i++ {
		if i != nums[i]-1 {
			return i + 1
		}
	}
	return n + 1
}
