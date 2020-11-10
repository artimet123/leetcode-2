package main

import "fmt"

func main() {
    channel := make(chan string, 2)

    fmt.Println("1")
    channel <- "h1"
    fmt.Println("...")
    msg1 := <-channel
    fmt.Println(msg1)
}