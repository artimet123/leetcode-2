package main 

import (
	"fmt"
)

// 归并排序
// 拆分
// 合并
func main(){

	var i =[]int{9,2,5,4,10}
	new := mergeSort(i)
	fmt.Println("new=", new)
}

func mergeSort(arr []int)[]int{
	fmt.Println("arr=", arr)
	if len(arr) <= 1 {
		return arr
	}

	// 取中间数
	var m = len(arr) / 2

	// 左边的数组不断的拆分
	b := mergeSort(arr[:m])
	fmt.Println("b=", b)

	// 右边的不断拆分
	c := mergeSort(arr[m:])
	fmt.Println("c=", c)

	new := merge(b, c)
	return new
}

func merge(left, right []int)[]int{
	var (
		new []int
	)

	var i, j int = 0, 0

	for i < len(left) && j < len(right) {
		// 谁比较小，将它放进数组，继续下一个
		if left[i] < right[j] {
			new = append(new, left[i])
			i++
		} else {
			new = append(new, right[j])
			j++
		}
	}

		//如果左边的数组还没比较完，右边的数都已经完了，那么将左边的数抄到大数组中(剩下的都是大数字)
		for i < len(left) {
			new = append(new, left[i])
			i++
		}
		//如果右边的数组还没比较完，左边的数都已经完了，那么将右边的数抄到大数组中(剩下的都是大数字)
		for j < len(right) {
			new = append(new, right[j])
			j++
		}
		fmt.Println("merge new=" , new)
	return new
}