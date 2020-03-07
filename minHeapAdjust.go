package main 

import "fmt"

func main(){
	var array = []int{1,3,5,2,6,8,9,12,34,56,104,25}
	heapSort(array)
	fmt.Println("array=", array)
	return
}

func heapSort(array []int){
	// 三步走
	// 1: 构造最小堆
	// 从最下面的非叶子结点开始
	for i := len(array)/2 - 1; i >= 0; i-- {
		minAdjustHeap(array, i, len(array))
	}

	// 交换和再次调整
	for i := len(array) - 1; i >= 0; i--{
		swap(array, 0, i)
		minAdjustHeap(array, 0, i)
	}
	return	
}

func minAdjustHeap(array []int, i, length int){
	// 先把第一个值保存起来
	var temp = array[i]

	for k := i * 2 +1; k < length; k = 2*k +1 {

		// 找出指定的那个非叶子结点的左右孩子的最小值，交换
		if k +1 < length && array[k] > array[k+1] {
			// 找到最小的
			k++
		}

		if array[k] < temp {
			array[i] = array[k]
			i = k
		} else {
			break
		}
	}
	array[i] = temp
	return
}

func swap(array []int, i, j int){
	var temp = array[j]
	array[j] = array[i]
	array[i] = temp
	return
}