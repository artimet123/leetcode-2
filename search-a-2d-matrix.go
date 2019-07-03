package main

import (
	"fmt"
)

func main() {
	m := [][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}
	b := martixSearch(m, 3)
	fmt.Println("b=", b)

}

func martixSearch(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])
	start, mid, end := 0, 0, m*n-1

	for start <= end {
		mid = start + (end-start)/2
		v := matrix[mid/n][mid%n]

		if target < v {
			end = mid - 1
		} else if target > v {
			start = mid + 1
		} else {
			return true
		}
	}

	return false
}
