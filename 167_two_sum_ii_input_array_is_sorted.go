package main 

import "fmt"

// 一个有序的数组，找出两者之和为target的序列的index的值,这个里面只要有一对符合就返回
func main(){

	var num =[]int{1,2,3,4}
	var target = 5
	index := twoSum(num, target)
	fmt.Println("index=", index)
}

func twoSum(numbers []int, target int)[]int{
	for i, n := range numbers{
		for j := i +1; j < len(numbers); j++ {
			sum := n+ numbers[j]
			if sum == target {
				return []int{i+1, j+1}
			}

			if sum > target {
				break
			}
		}
	}
	return nil
}