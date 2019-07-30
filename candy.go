package main 

import "fmt"

func main(){

	r := []int{1,3,4,2,4,6,9,5,6,8,15}
	ret := candy(r)
	fmt.Println("ret=", ret)
}

func candy(ratings []int) int{
	n := len(ratings)
	if n == 0 {
		return 0
	}

	candys := make([]int, n) // 这个必须是长度固定的，如果不是，会造成panic，默认会填充全是0
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] && candys[i] <= candys[i-1] {
			candys[i] = candys[i-1] + 1
		}
	}
	fmt.Println("first-candys=",candys)
	for i := n - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] && candys[i-1] <= candys[i] {
			candys[i-1] = candys[i] + 1
		}
	}
	fmt.Println("second-candys=",candys)
	for _, c := range candys {
		n += c
	}
	return n

}