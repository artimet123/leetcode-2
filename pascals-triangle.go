package main

import "fmt"

func main() {
	var r int = 5
	ret := generate(r)
	fmt.Println("ret=", ret)

}

func generate(rows int) [][]int {
	var row []int
	var ret [][]int

	for i := 1; i <= rows; i++ {
		for j := i - 1; j >= 2; j-- {
			row[j-1] += row[j-2] // 每行的计算，是从右往左计算
		}

		row = append(row, 1)
		val := make([]int, len(row))
		copy(val, row)
		ret = append(ret, val)
	}
	return ret
}
