package main 

import "fmt"

// 创建工厂结构体并实现工厂接口
// 简单工厂模式：工厂太单一了，唯一了。所以在工厂方法中，我们把工厂当成一个接口，不同的工厂，可以造就不同的商品，解耦。
// 通过实现工厂接口，可以创建多种工厂。

// 感觉就是：产品实现产品接口，工厂实现工厂接口。产品还是统一从工厂出货，只是工厂出货不再是个具体的产品，而是一个接口
type FactoryInterface interface {
	CreateProduct(t string)ProductInterface
}

// 创建工厂结构体
type Factory1 struct {}

func (f Factory1)CreateProduct(t string)ProductInterface{
	switch t{
	case "product1":
		return Product1{}
	default:
		return nil
	}
}

// extra扩展第二家工厂，只要实现接口
type Factory2 struct{
	Factory1
}

// extra 虽然继承，但是可以多态
func (f Factory2)CreateProduct(t string)ProductInterface{
	switch t{
	case "product2":
		return Product1{}
	default:
		return nil
	}
}

// 产品接口
type ProductInterface interface{
	Intro()
}

// 创建产品1并实现产品接口
type Product1 struct{}

type Product2 struct{}

func(p Product1)Intro(){
	fmt.Println("this is product1")
}

func(p Product2)Intro(){
	fmt.Println("this is product2")
}

func main(){
	f := new(Factory1)

	p := f.CreateProduct("product1")

	p.Intro()

	f2 := new(Factory2)
	p2 := f2.CreateProduct("product2")
	p2.Intro()
}
