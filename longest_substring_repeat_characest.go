package main 

import "fmt"

func main(){
	var s = "123355411"
	 c := lengthOfLongestSubstring(s)
	 fmt.Println("logest charcest = ", c)
}

// 计算不重复值的最长长度
func lengthOfLongestSubstring(s string) int {
	var max, start int
	letter := make(map[rune]int)
	for i := 0; i < len(s); i++ {
		c := rune(s[i])
		prev, ok := letter[c]
		if ok && prev >= start {
			// start 代表的是不重复开始的位置，从1开始计数，如果出现重复，以最后的为准
			// prev 指向的是与该值一致的最新的位置
			start = prev + 1
		}

		letter[c] = i // 影响了prev的最新值
		len := i - start + 1
		if max < len {
			max = len
		}
	}
	return max
}