package main 

import "fmt"

func main(){
	var nums = []int{1,2,4,6,6,78}
	ok := containDuplicate(nums)
	fmt.Println("ok=", ok)

}

func containDuplicate(nums []int)bool{
	m := make (map[int]struct{}, len(nums))
	for _, n := range nums{
		if _, ok := m[n]; ok {
			return true
		}
		m[n] = struct{}{}
	}
	return false
}