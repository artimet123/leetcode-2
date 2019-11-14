package main

import "fmt"
/*
一个数是不是快乐是这么定义的：对于一个正整数，每一次将该数替换为他每个位置上的数字的平方和，
然后重复这个过程直到这个数变为1，或是无限循环但始终变不到1。
如果可以变为1，那么这个数就是快乐数。

样例：19 就是一个快乐数：19->82->68->100->1
*/

func main() {
	var n = 2

	ok := isHappy(n, nil)
	fmt.Println("ok=", ok)

}

func isHappy(n int, t []int) bool {
	if n == 1 {
		return true
	}
	fmt.Println("n=", n, " t=", t)
	t = append(t, n) // 记住循环的数字，防止无限循环。毕竟在数组里面的数，说明已经出现过计算，没必要继续计算下去
	var nn int

	for n > 0 {
		i := n % 10
		nn += i * i
		n /= 10
	}

	for _, v := range t {
		if v == nn {
			return false
		}
	}
	return isHappy(nn, t)
}
