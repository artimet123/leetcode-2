package main 

import "fmt"

// 抽象工厂模式
// 多个工厂实现各自的接口
// 多个产品实现各自的接口
// 每一个工厂可以生产各自的产品线
type FactoryInterface interface{
	CreatePigBuns()ProductInterface // 创建猪肉产品
	Create3SBus()ProductInterface // 创建三鲜产品
}

type ProductInterface interface{
	Intro()
}

type BJPigBuns struct{}

func (b BJPigBuns)Intro(){
	fmt.Println("北京猪肉包子")
}

type BJ3SBuns struct{}

func (b3 BJ3SBuns)Intro(){
	fmt.Println("北京三鲜包子")
}

type TJPigBuns struct{}

func (b TJPigBuns)Intro(){
	fmt.Println("天津猪肉包子")
}

type TJ3SBuns struct{}

func (b3 TJ3SBuns)Intro(){
	fmt.Println("天津三鲜包子")
}

// 因为工厂抽象为的是接口

type BJFactory struct{}

type TJFactory struct{}

// 各自的实现方法
func (bj BJFactory)CreatePigBuns()ProductInterface{
	return BJPigBuns{}
}

func (bj BJFactory)Create3SBus()ProductInterface{
	return BJ3SBuns{}
}

func (tj TJFactory)CreatePigBuns()ProductInterface{
	return TJPigBuns{}
}

func (tj TJFactory)Create3SBus()ProductInterface{
	return TJ3SBuns{}
}



func main(){
	var f FactoryInterface
	f = new(BJFactory) // 因为工厂已经实现了接口，所以可以用结构体变量赋值给接口变量
	b := f.Create3SBus() // 其实b也是接口，只是指向了BJ3SBuns结构体
	b.Intro() // 结构体对应的方法
}