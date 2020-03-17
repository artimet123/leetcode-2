package main 

import "fmt"

func main(){
	// 采用压栈的方式计算
	var s string = "[(())]{}"

	b := isMatch(s)
	fmt.Println(b)

	var s2 string = "][(())]{"

	b2 := isMatch(s2)
	fmt.Println(b2)

}

func isMatch(s string)bool{
	var m = make(map[byte]byte)
	// 初始化
	m['('] = ')'
	m['{'] = '}'
	m['['] = ']'

	sArray := []byte{}
	if len(s) == 0 {
		return true
	}

	if len(s) / 2 == 1{
		return false
	}

	for i, b := range s {
		if i == 0 && (b == ')' || b == ']' || b == '}'){
			return false
		}

		if b == '[' || b == '{' || b == '('{
			sArray = append(sArray, byte(b))
		}else {
			// 拿出最后一个元素和字符比较
			left := sArray[len(sArray) - 1]
			fmt.Println("left=", string(left))
			if m[left] != byte(b) {
				return false
			}
			sArray = sArray[:len(sArray) - 1]
		}
	}

	if len(sArray) != 0 {
		return false
	}

	return true
}