package main

import "fmt"

/**
Go接口是一种特殊的数据类型
这里定义一个汽车接口，然后实现“驾驶、倒车”方法
*/

type Car interface {
	Drive()
	Back()
}

// 定义丰田汽车结构体，并实现Car接口
type ToyotaCar struct{}

/**
实现Car接口的Drive方法
*/
func (toyotaCar ToyotaCar) Drive() {
	fmt.Println("我是丰田车，正在驾驶。。。")
}

/**
实现Car接口的Back方法
*/
func (toyotaCar ToyotaCar) Back() {
	fmt.Println("我是丰田车，正在倒车。。。")
}

type BMWCar struct{}

func (bmwCar BMWCar) Drive() {
	fmt.Println("我是宝马车，正在驾驶。。。")
}

func (bmwCar BMWCar) Back() {
	fmt.Println("我是宝马车，正在倒车。。。")
}

func (bmwCar BMWCar) Park() {
	fmt.Println("我是宝马车，正在停车。。。")
}

func main() {
	// 声明Car接口的实例，并通过强转指向ToyotaCar结构体的实例
	var car Car = new(ToyotaCar)
	/*
		我是丰田车，正在驾驶。。。
		我是丰田车，正在倒车。。。
	*/
	car.Drive()
	car.Back()

	/**
	宝马车多了一个方法Park，看通过Car接口的实例能否访问
	*/
	var bmwCar Car = new(BMWCar)
	/*
		我是宝马车，正在驾驶。。。
		我是宝马车，正在倒车。。。
	*/
	bmwCar.Drive()
	bmwCar.Back()
	// bmwCar.Park()//bmwCar.Park undefined (type Car has no field or method Park)
}
