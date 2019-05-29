package main 

import (
	"fmt"
)

// 装饰器特别适合重复操作同一类型的操作，并且从代码上减少变量使用错误
func main (){
	bFunc := uniqGetDealerId()

	bFunc("a")
	bFunc("b")
	bFunc("a")
	bFunc("c")

}

func uniqGetDealerId() func(string){
	idMap := make(map[string]string)

	return func(dealerId string) {
		if v, ok := idMap[dealerId]; ok {
			fmt.Println("this dealerid=", v)
			return
		}

		// key和value是一个值
		idMap[dealerId] = dealerId
		return
	}
}