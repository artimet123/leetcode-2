package main

import (
	"fmt"
)

func main(){

}

func insertionSort(arr []int) []int {
	// 因为前面是排好的序，所以后面只需要找到不大于current的那个数字替换，但是多减了个1，所以最后还得再+1去替换
    for i := range arr {
        preIndex := i - 1
        current := arr[i]
        for preIndex >= 0 && arr[preIndex] > current {
            arr[preIndex+1] = arr[preIndex]
            preIndex -= 1
        }
        arr[preIndex+1] = current
    }
    return arr
}