package main

import (
	"fmt"
)

func main(){
	var s = "ab"
	ret := longestPalindrome(s)
	fmt.Println(ret)
}

func longestPalindrome(s string) string {
    var maxLen int
    var maxStr string
    for i := 0; i < len(s); i++ {
        for j := i + 1; j < len(s)+1; j++ {
            len := isPalindrome(s[i:j])
            if len > maxLen {
                maxStr = s[i:j]
                maxLen = len
            }
        }
    }
    return maxStr
}

func isPalindrome(s string) int {
    var newStr string
    for i := len(s)-1; i >= 0; i--{
        newStr += s[i:i+1] // 将字符从后往前重新组装
    }
    if s == newStr {
        return len(s)
    }
    return 0
}