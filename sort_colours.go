package main 

import (
	"fmt"
)

/* 排序0，1，2三种元素*/
func main(){
	c1 := []int{1,2,1,0,2}
	sortColors(c1)
	fmt.Println("c1=", c1)

}

func sortColors(nums []int) {
	l, r := 0, len(nums)-1
	for i := l; i <= r; i++ {
		for i < r && nums[i] > 1 {
			// 该循环不会改变i的值，只是会改变i所在位置的数组的值；但是会改变r的值，i和r所在的位置的指向的值会发生颠倒
			// 不符合for的条件的，直接跳出循环，执行下面的那个循环
			// 作用：保证只循环一遍，就可以做到0，1，2的排序
			nums[i], nums[r] = nums[r], nums[i]
			r--
		}
		for i > l && nums[i] < 1 {
			nums[l], nums[i] = nums[i], nums[l]
			l++
		}
	}
}