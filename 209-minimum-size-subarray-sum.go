package main 

import (
	"fmt"
)

func main(){
	var s = 7
	nums := []int{1,2,3,4,5,6,3}
	n := minSubArrayLen(s, nums)
	fmt.Println("n= ", n)

}

func minSubArrayLen(s int, nums []int)int{
	var i,j, t,n int

	for j < len(nums){
		// 先让数组一直加
		for t < s && j < len(nums){
			t += nums[j]
			j++
		}

		// 再从头扣减
		for t >= s && i < len(nums) {
			if  n == 0 || n > j - 1 {
				n = j -1
			}
			t -= nums[i]
			i++
		}
	}
	return n
}