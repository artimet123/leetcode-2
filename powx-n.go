package main

import (
	"fmt"
)

func main() {
	var (
		x float64 = 3.2
		n int
	)

	n = -2

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
	} else {
		ret = myPow(x, n/2)
		ret *= ret * x
	}
	return ret

}