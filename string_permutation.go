package main

import (
	"fmt"
)

func main(){
	var a = []string{"a", "b", "c"}
	permutation(a, 0, len(a))
	fmt.Println(a)


}

func permutation(a []string, index, length int){
	if index == length {
		var s string
		for _, v := range a {
			s += v 
		}
		fmt.Println("s-mao=",s)
	} else {
		for i := index; i < length; i++ {
			if i != index && a[i] == a[index]{
				continue
			}
			swap(a, i, index)
			permutation(a, index +1, length)
			swap(a, i, index)
		}
	}
}

func swap(a []string, i, j int){
	var temp = a[i]
	a[i] = a[j]
	a[j] = temp
}