package main

import (
	"fmt"
)

func main(){
	var s = "abcdeedcb"
	ret := longestPalindrome(s)
	fmt.Println(ret)

}

func longestPalindrome(s string) string {
    if len(s) < 2 {
        return s
    }
    // P[i][j]表示回文，P[i+1][j-1]也是回文
    longest := s[0:1] // 默认是第一个
    for i := 1; i < len(s); i++ {
        for rightStep := 0; rightStep < 2; rightStep++ {
        	// 这个地方之所以会是两个循环的原因是：回文是奇数还是偶数

            for p, q := i-1, i+rightStep; p >= 0 && q < len(s) && s[p] == s[q]; {
                if q-p+1 > len(longest) {
                    longest = s[p : q+1]
                }
                p--
                q++
            }

        }
    }
    return longest
}
