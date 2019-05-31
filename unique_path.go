package main

import (
	"fmt"
)

// 计算头走到尾的路径，要求不一样。
// 不太好理解，坐标的+，其中坐标为同一行的前一个 + 同一列的前一个

// 这个算法同样适用于在某个坐标设置障碍物，方法就是定义好哪些位置是障碍物比如
// if i == 2 && j == 2 设置为0

func main() {
	m, n := 4, 3
	l := uniquePaths(m, n)
	fmt.Println("l=", l)
}
 
func uniquePaths(m int, n int) int {
	nums := make([]int, n)
	nums[0] = 1
for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			nums[j] += nums[j-1]
		}
	}
	return nums[n-1]
}
