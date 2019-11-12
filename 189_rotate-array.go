package main

import (
	"fmt"
)

/*
思路：把原数组划分为两个部分来看：前 n - k 个元素 [1,2,3,4]和 后k个元素 [5,6,7]，进行分开处理

    定义 reverse 逆转方法：将数组元素反转，比如 [1,2,3,4] 逆转后变成 [4,3,2,1]
    对前 n - k 个元素 [1,2,3,4] 进行逆转后得到 [4,3,2,1]
    对后k个元素 [5,6,7] 进行逆转后得到 [7,6,5]
    将前后元素 [4,3,2,1,7,6,5] 逆转得到：[5,6,7,1,2,3,4]
    注意：还要处理 k > 数组长度的情况，对 k 进行取模
*/
func main() {
	nums := []int{2, 4, 5, 1, 5, 7, 9}
	rotate(nums, 4)
	fmt.Println("nums=", nums)
}

func rotate(nums []int, k int) {
	size := len(nums)

	if k %= size; k > 0 && k < size {
		// 一次拆分
		rorateArray(nums[:size-k]) // 前半部分
		rorateArray(nums[size-k:]) // 后半部分
		rorateArray(nums)          // 全部
	}
}

func rorateArray(nums []int) {
	for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}
