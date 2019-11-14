package main

import "fmt"

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
