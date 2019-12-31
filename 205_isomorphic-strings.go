package main 

import "fmt"

func main(){
	s, t := "egg", "ftt"
	ok := isIsomorphic(s, t)
	fmt.Println("ok=", ok)

}

func isIsomorphic(s, t string)bool {
	if len(s) != len(t){
		return false
	}

	// 定义两个hashmap
	ss, tt := make(map[byte]int), make(map[byte]int)

	for i, j := 0, 0; i < len(s); i,j = i+1, j+1{
		b1, b2 := s[i], s[j]
		if ss[b1] != tt[b2]{
			return false
		}
		ss[b1], tt[b2] = i+1, j+1
	}
	return true
}