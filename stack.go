package main 

import "fmt"

type stack struct{
	Array []string
	Size int
	Top int
}

func NewStack(size int)*stack{
	return &stack{
		Array: make([]string, size),
		Size: size,
		Top: -1,
	}
}

func (s *stack)IsEmpty()bool{
	return s.Top == -1
}

func (s *stack)IsFull()bool{
	return s.Top == s.Size - 1
}

func(s *stack)Push(item string)bool{
	if s.IsFull(){
		return false
	}

	s.Top++
	s.Array[s.Top] = item
	return true
}

func(s *stack)Pop()(string, bool){
	if s.IsEmpty(){
		return "", false
	}

	item := s.Array[s.Top]
	s.Top--
	s.Array[s.Top] = item
	return item, true
}

func main(){
	stack1 := NewStack(10)
	stack1.Push("mao")
	stack1.Push("ming")

	fmt.Println(stack1.Array)

}