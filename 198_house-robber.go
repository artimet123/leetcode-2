package main 

import (
	"fmt"
)	

func main(){
	var nums =[]int{1,3,18,21,7,8,10}
	ret := robHouse(nums)
	fmt.Println("ret=", ret)
}

func robHouse(nums []int)int{
	if len(nums) == 0 {
		return 0
	}

	vals := make([]int, len(nums))
	for i, num := range nums {
		if i == 0 {
			vals[i] = num
		} else if i == 1 {
			vals[i] = max(num, vals[i-1])
		} else {
			vals[i] = max(num+vals[i - 2],vals[i - 1])
		}
	}
	return vals[len(nums) - 1]
}

func max(a, b int)int{
	if a == b {
		return a
	} else if a < b {
		return b
	} else {
		return a
	}
	return 0
}