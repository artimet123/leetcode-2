package main

import "fmt"

func main() {
	var n int = 12
	res := reverseBits(n)
	fmt.Println("res=", res)
}

func reverseBits(n int) (res int) {
	var (
		flag int = 1
	)

	for i := 0; i < 8; i++ {
		//这个循环都是好次数，在一个10进制转换为2进制前十不知道占多少位的
		res = res*2 + (flag & n) // 先扩大两倍，再加也是可以的。
		n = n >> 1
	}
	return res
}
