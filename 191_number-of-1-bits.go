package main

import "fmt"

func main() {

	var n uint32 = 15
	ret := bit1Count(n)
	fmt.Println("ret=", ret)

}

func bit1Count(n uint32) uint32 {
	var c uint32 = 0
	for n > 0 {
		c += (n & 0x01)
		n >>= 1
	}
	return c
}
