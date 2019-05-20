package main 

import "fmt"

func main(){
	var (
		numRows int
		s = "MAOYONGMING"
	)
	numRows = 4
	s = convert(s, numRows)
	fmt.Println("s=", s)

}
func convert(s string, numRows int) string {
	i, size := 0, len(s)
	ss := make([]string, numRows)
	for i < size {
		for j := 0; j < numRows && i < size; j++ {
			ss[j] += string(s[i])
			i++
		}

		// 最后一列和第一列都不会append值，只有中间部分从后往前置
		for j := numRows - 2; j > 0 && i < size; j-- {
			ss[j] += string(s[i])
			i++
		}
	}
	var ret string
	for j := 0; j < numRows; j++ {
		ret += ss[j]
	}
	return ret
}
