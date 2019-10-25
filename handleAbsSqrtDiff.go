package main

// 计算有序数组中，平方值不同的 和
import (
	"fmt"
	"math"
)

func main() {
	var nums = []float64{-10, -7, -6, -2, 0, 1, 2, 3, 5, 7, 17}
	ret := handle(nums)
	fmt.Println("ret=", ret)
}

func handle(nums []float64) (count int) {

	if len(nums) == 0 {
		return
	}

	// 定义两个指针(说白了就是计数)
	var (
		i, j = 0, len(nums) - 1
	)

	for i <= j {
		num1 := math.Abs(nums[i])
		num2 := math.Abs(nums[j])
		fmt.Println("i=", i, " j=", j)
		fmt.Println("count=", count)
		if num1 > num2 {
			count++
			for i <= j && math.Abs(nums[i]) == num1 {
				i++
			}

		} else if num1 < num2 {
			count++
			for i <= j && math.Abs(nums[j]) == num2 {
				j--
			}
		} else {
			count++
			for i <= j && math.Abs(nums[i]) == num1 {
				i++
			}
			for i <= j && math.Abs(nums[j]) == num2 {
				j--
			}
		}
	}
	return
}
