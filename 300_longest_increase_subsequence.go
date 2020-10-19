package main

import "fmt"

// 找出最长的上升的子序列
// 其中的那个替换A是，拿元素和数组中已存在的 > A的最小元素，去替换;没有比A大的就append A。
func main(){
	var nums []int = []int{10,11,12,13,14,16,1,2,3,5,7,15}
	len := LengthOfLIS(nums)
	fmt.Println(len)


}

func LengthOfLIS(nums []int)int{
	n := len(nums)
	if n == 0 {
		return 0
	}

	var dp []int = make([]int, n) // 创建dp数组，用来存放最后的结果
	var m int = 0

	for _, num := range nums {
		var low, high int
		high = m

		for low < high {
			mid := low + (high - low) / 2
			fmt.Println(mid, num)

			if dp[mid] < num {
				low = mid + 1
				fmt.Println("11111-", low, num)
			} else {
				fmt.Println("22222-", high, num)
				high = mid
			}
		}
		fmt.Println("33333-", low, num)
		dp[low] = num
		if low == m {
			m++
		}

	}
	return m
}

