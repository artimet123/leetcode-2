package main 
import (
	"fmt"
)

func main(){
	var  a = []int{1,2,3,4,5,6,7,9,10,13,14,20}
	sort(a)
	fmt.Println("a=", a)
}

func sort(arr []int){
	var (
		i,j int
		length = len(arr)
	)

	// 第一步 构造堆 时间复杂度为o(n)
	for i = length/2 - 1; i >= 0; i--{
		adjustHeap(arr, i, length)
	}

	// 第二步 顶端数据和最右数据交换，并重新构造堆(要注意：不要包含最后的一个数据，length - 1说明了一切) 时间复杂度为o(nlogN)
	for j = length - 1; j > 0; j-- {
		swap(arr, 0, j)
		adjustHeap(arr, 0, j)
	}
}

func adjustHeap(a []int, i, length int){
	var (
		temp int = a[i]
		k int
	)

	for k = i * 2 +1; k < length; k = k*2 + 1 {

		if (k + 1 < length && a[k] < a[k+1]){
			k++
		}

		if a[k] > temp {
			a[i] = a[k]
			i = k
		} else{
			break
		}	
	}
	a[i] = temp
}

func swap(arr []int, a, b int) {
	var temp int = arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}
