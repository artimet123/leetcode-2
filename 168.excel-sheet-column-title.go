package main 

import "fmt"

// 给出一个数字，确定好是在excel列的哪个里面
func main(){
	var n = 28
	s := convertToTitle(n)
	fmt.Println("s=", s)

}

func convertToTitle(n int)string{
	const l = "ZABCDEFGHIJKLMNOPQRSTUVWXY"
	var s string
	for n > 0 {
		i := n % 26
		n = (n - 1) / 26 // 不明白这个里面为什么要剪去1
		s = string(l[i]) + s

	}
	return s
}