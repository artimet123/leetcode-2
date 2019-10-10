package main

import (
	"fmt"
	"strings"
)

func main(){

	s := "Hello World! golang 世界"

	new := reverseWords(s)

	fmt.Println("new=", new)

}


func reverseWords(s string) string {
	arr := strings.Split(s, " ") // 分隔字符串
	res := []string{}
	for i := 0; i < len(arr); i++ {
		if arr[i] != "" {
			res = append(res, arr[i])
		}
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return strings.Replace(strings.Trim(fmt.Sprint(res), "[]"), " ", " ", -1)

}