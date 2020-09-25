package main 

import (
	"fmt"
)

func main(){
	var s = []int{10,9,4,3,2,1}
	max := MaxFunc(s)
	fmt.Println(max)

}

func MaxFunc(a []int)(max int) {
	if len(a) == 0 {
		return
	}

	if len(a) == 2 {
		max = a[1] - a[0]
		return
	}

	max = a[1] - a[0]
	var (
		start = Min(a[0], a[1]) // 找出最小值作为起点
	)

	for i := 2; i < len(a); i++{
		max = Max(max, a[i] - start)
		start = Min(start, a[i])
	}
    return
}

func Min(a, b int)int{
	if a <= b {
		return a
	}
	return b
}

func Max(a, b int)int{
	if a <= b {
		return b
	}
	return a
}