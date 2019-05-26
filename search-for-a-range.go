package main

import "fmt"

func main() {
	var (
		noNew  []int
		nums       = []int{1, 2, 3, 4, 5, 6, 6, 8, 9, 14}
		target int = 5
	)
	noNew = searchRange(nums, target)
	fmt.Println("num-new=", noNew)
}

func searchRange(nums []int, target int) []int {
	if nums == nil {
		return []int{-1, -1}
	}
	// r = 9
	l, r := 0, len(nums)-1

	for nums[l] < nums[r] {
		m := (l + r) / 2
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			if nums[l] == nums[m] {
				r--
			} else {
				l++
			}
		}
	}
	if nums[l] == nums[r] && nums[l] == target {
		return []int{l, r}
	}
	return []int{-1, -1}
}
