package main 

import (

	"fmt"
)

// 简单工厂模式
// 感觉就好像工厂是固定的，根据不同的需求创造出不同的产品(同一层级),再根据产品的多态计算出各个产品的特性

// 大概要做的事情就是，当我们想要创建一个对象的时候，调用同一个方法，传入不同的参数就可以返回给我们不同的对象了
//当然，前提是这些对象对应的类都实现了相同的接口

type Product interface {
	create()
}

// 工厂类
type Factory struct {}

func (c Factory)Generate(name string) Product {
	switch name{
	case "product1":
		return Product1{}
	case "produtct2":
		return Product2{}

	default:
		return nil
	}
}

// 产品1实现接口
type Product1 struct{}
func (p1 Product1)create(){
	fmt.Println("Hello world p1")
}

type Product2 struct{}

func (p2 Product2)create(){
	fmt.Println("Hello World p2")
}

// 扩展新问题，继承
type Product3 struct {
	Product2
}

func (p3 Product3)create(){
	fmt.Println("Hello World p3")
}

func main(){
	factory := new(Factory) // 工厂声明

	// 获取不同的实例
	// 因为p1和p2实现了函数create，也就是说实现了接口，那这个时候其实实例是可以传给interface的
	p1 := factory.Generate("product1")
	p1.create()

	p2 := factory.Generate("produtct2")
	p2.create()

	var p3 Product3
	p3.create()

}