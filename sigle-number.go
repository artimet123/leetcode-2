package main 

import "fmt"

func main(){
	var a = []int{1,1,2,3,3}
	_ = singleNumber(a)
}

func singleNumber(nums []int)int{
	var x int

	for _, v := range nums{
		x = x ^ v
		fmt.Println("x=", x)
	}
	return x
}