package main

import "fmt"

func main(){
	var nums []int = []int{1,3,-1,-3,5,3,6,7}
	var k int = 3
	max := maxSlidingWindow(nums, k)
	fmt.Println(max)

}

func maxSlidingWindow(nums []int, k int)(max []int){
	if k == 0 {
		return 
	}

	var (
		window []int = make([]int, k) // nums数组的下标
	)

	for i := 0; i < len(nums); i++ {
		// 当第一个元素的下标 + wind的长度k, 已经超出循环的位置时;说明第一个元素不应该出现在新的比较窗口了
		if i >= k && window[0] <= i - k {
			window = window[1:] // 第一个元素出去
		}
		
		for len(window) > 0 && nums[window[len(window) - 1]] <= nums[i] {
			// 从wind数组的最右边开始寻找 比 当前大数组的值小的，直接删除，这样才能保留数组中最左边是最大的。
			window = window[:len(window) - 1]
		}
		window = append(window, i) // 当前的元素append进去，保证了当前窗口的数组，最左边为最大

		if i >= k - 1 {
			// 因为是达到3个后，每往后移动一个都是一个窗口
			max = append(max, nums[window[0]])
		}
	}
	return
}