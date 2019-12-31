package main 

// 找出第K个元素，是寻找：全部数组中排名第K个的最大元素
import (
	"fmt"
)

func main(){
	var k = 5
	var nums = []int{1,3,4,6,8,3,5,8,10,14}
	ret := findKthLargest(nums, k)
	fmt.Println("ret=", ret)
	fmt.Println("nums=", nums)

}

// 这样会导致从大到小排序
func findKthLargest(nums []int, k int)int{
	num := nums[0] // 第一个元素

	l , r := 0, len(nums) - 1 // 记录第一个和最后一个位置

	// 这样一轮下来，左边全是小num的值，右边全是大于num的值
	for l < r {
		for l < r && nums[r] <= num {
			r--
		}
		nums[l] = nums[r] // 这样r的位置就可以认为是空坑
		for l < r && nums[l] >= num {
			l++
		}
		nums[r] = nums[l] // 填满了r位置，l位置出现了新坑
	}
	nums[l] = num // 最后要把l填上基准值

	if l+1 == k {
		return num // 不需要找了，直接就是
	} else if l +1 > k {
		return findKthLargest(nums[:l+1], k) // 左半部查找
	} else {
		return findKthLargest(nums[l+1:], k - l -1) // 右半部 查找
	}


}