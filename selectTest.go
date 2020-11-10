package main

import (
    "fmt"
)
//定义几个变量，其中chs和numbers分别代表了包含了有限元素的通道列表和整数列表
    var ch1 = make(chan int, 3)
    var ch2 = make(chan int, 2)
var chs = []chan int{ch1, ch2}
var numbers = []int{1,2,3,4,5}

func main(){
    var i int
   // var ok bool

    select {
        case <- ch1:
           fmt.Println("1th case is selected.")
        case getChan(1) <- getNumber(3):
            fmt.Println("2th case is selected.")
        default:
            fmt.Println("default!.")
    }

    i, _  = <- chs[1]
    fmt.Println(i)
}

func getNumber(i int) int {
    fmt.Printf("numbers[%d]\n", i)
    return numbers[i]
}

func getChan(i int) chan int {
    fmt.Printf("chs[%d]\n", i)
    return chs[i]
}