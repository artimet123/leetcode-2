package main

import "fmt"

/*计算杨辉三角中的第N行的数组，其中要注意的是，行是从0开始计数*/
func main() {
	var rowIndex = 4
	a := getRow(rowIndex)
	fmt.Println("a=", a)
}

// 采用的方法是：先补齐后，再去修改其中需要修改的值。
func getRow(row int) []int {
	if row == 0 {
		return []int{1}
	}

	rowArray := append(getRow(row-1), 1)

	for i := row - 1; i > 0; i-- {
		rowArray[i] += rowArray[i-1]
	}
	return rowArray
}
