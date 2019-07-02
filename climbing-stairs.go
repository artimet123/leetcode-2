package main

import (
	"fmt"
)

/*
	爬台阶问题，每一次可以爬一个或者两个台阶，问一个有n个台阶，一共有多少种方法爬到n
	n = 1时，1种;
	n = 2时，2种;
	n>2时，对于每一个台阶i，要到达台阶，最后一步都有两种方法，从i-1迈一步，或从i-2迈两步。

　　也就是说到达台阶i的方法数=达台阶i-1的方法数+达台阶i-2的方法数
*/
func main() {
	var n = 6
	c := climbStairs(n)
	fmt.Println("c=", c)

}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	n1, n2 := 1, 2

	for i := 3; i <= n; i++ {
		n2, n1 = n1+n2, n2
	}
	return n2
}
