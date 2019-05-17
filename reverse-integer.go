package main

import "fmt"

func main(){
	var x int = 12345
	fmt.Println("final data=", reverse(x))
}

func reverse(x int) int {
	var neg bool
	if x >= 0 {
		neg = true
	} else {
		x *= -1
	}
	var y int
	for x > 0 {
		y = y*10 + x%10
		x = x / 10
		if y > 2147483647 {
			return 0
		}
	}
	if neg == false {
		y *= -1
	}
	return y
}
