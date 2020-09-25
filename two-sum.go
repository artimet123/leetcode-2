package main

import "fmt"

// 给定一个整数数组和一个目标值，找出数组中和为目标值的 两个数
// 时间复杂度o(n) = n
func main() {
	var (
		target int = 9
		nums       = []int{2, 7, 22, 13, 7}
		b      []int
	)
	b = twoSum(nums, target)
	fmt.Println("b=", b)
}

func twoSum(nums []int, target int) []int {
	is := make(map[int]int)
	for i, n := range nums {
		/* 两个值相等时，会造成循环的次数增大，因为下面的循环是从左开头开始查找
		if _, ok := is[n]; ok{
			continue
		}
		*/
		is[n] = i
	}

	for i, n := range nums {
		if j, ok := is[target-n]; ok && i != j {
			return []int{i, j}
		}
	}
	return nil
}
