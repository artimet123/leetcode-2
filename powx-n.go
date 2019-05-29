package main

import (
	"fmt"
)

func main() {
	var (
		x float64 = 2
		n int
	)

	n = 5

	c := myPow(x, n)
	fmt.Println("c=", c)
}

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}

	var ret float64
	if n < 0 {
		ret = 1/ myPow(x, -n)
	} else if n == 0 {
		ret = 1
	} else if n%2 == 0 {
		ret = myPow(x, n/2)
		ret *= ret // 正好除2 余0，说明 结果*结果 即可
	} else {
		ret = myPow(x, n/2)
		ret *= ret * x // 不能整除时，说明需要加个数字
	}
	return ret
}