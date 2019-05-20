package main 

import "fmt"

func main(){
	var (
		s = 12321
		ok bool
	)
	ok = isPalindrome(s)
	fmt.Println("ok=", ok)

}

func isPalindrome(x int) bool {
	x1, x2 := x, 0
	for x > 0 {
		n := x % 10
		x2 = x2*10 + n
		x = x / 10
	}
	return x1 == x2
}
