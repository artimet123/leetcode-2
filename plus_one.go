package main

import "fmt"

// 这个算法的目标是：把数中的每一位的数字都放在数组中。比如 13456345 -> [1,3,4,5,6,3,4,5]
// 再 +1即可
func main() {

	var digits = []int{1, 3, 4, 5, 6, 3, 4, 5}
	digits = plusOne(digits)

	fmt.Println(" digits=", digits)

}

func plusOne(digits []int) []int {
	c := 1

	for i := len(digits) - 1; i >= 0; i-- {
		n := digits[i] + c
		digits[i] = n % 10
		c = n / 10
	}

	// 最高位 /10 余数 为1，说明，需要向前进一位
	if c == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
