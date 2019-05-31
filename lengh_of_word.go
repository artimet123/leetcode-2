package main

import (
	"fmt"
)

type Interval struct {
	Start int
	End   int
}

func main() {
	var s  = "Hello World  "
	l := lengthOfLastWord(s)
	fmt.Println("l=", l)
}

func lengthOfLastWord(s string) int {
	// 计算最后一个单词的长度，如果最后一个是一个或多个连续的空格，也还是最后一个不为空的词
	var n1, n2 int

	for _, c := range s{
		if c != 0x20 {
			n2++
		} else if c == 0x20 && n2 != 0 {
			n1, n2 = n2, 0
		}
	}

	if n2 == 0 {
		n2 = n1
	}

	return n2
}