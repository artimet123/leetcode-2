package main

import "fmt"

func main() {
	var (
		nums  = []int{3, 4, 5, 8, 9}
		value = 3
	)

	n := removeElement(nums, value)
	if n != 4 {
		fmt.Println("err")
	}
}

func removeElement(nums []int, v int) int {
	var n int
	for _, num := range nums {
		if num != v {
			nums[n] = num
			n++
		}
	}
	return n
}
