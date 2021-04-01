package main

import (
	"fmt"
	"reflect"
)

/**
Go语言中，可以通过reflect.ValueOf()来反射变量的值相关信息
*/

// 这个函数实现通过反射的方式获取变量的值信息，并构造出一个对应类型，对应值的变量
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	/**
	reflect.ValueOf(x)返回的是reflect.Value这个结构体实例
	这个实例中，包含了变量的值，我们可以转换为具体的对应的变量的值，还可以修改这个变量的值
	*/
	fmt.Printf("v的类型：%T, v的值%v\n", v, v)

	/**
	通过ValueOf来转换成对应的类型的变量，并与传入的值相同
	比如传入一个float32，值是1.112，我们这里的需求是得到一个float32类型，值是1.112的变量

	具体做法是：通过reflect.Value的内置方法来转换，比如v.Int()可以获取一个int64类型的值，但是事先需要知道
	值的类型，这个时候可以调用v.Kind()方法
	*/
	// reflect.Value 结构体也有Type()方法
	k := v.Kind()
	switch k {
	case reflect.Float32:
		ret := float32(v.Float()) // 返回float64，需要强制转化为float32
		fmt.Printf("得到的数据类型：%T，值：%v\n", ret, ret)
	case reflect.Int32:
		ret := int32(v.Int()) // 返回int64，需要强制转化为int32
		fmt.Printf("得到的数据类型：%T，值：%v\n", ret, ret)
	}
}

/**
这个函数的目标是，反射变量的值信息，并修改变量的值
比如传入一个float32类型的变量，那么将这个变量的值*10，并且在外部的变量也改变了值

由于Go语言中参数传递都是值拷贝，如果需要外部也感觉到传入变量的值改变了，需要传入对应变量的指针类型
*/
func reflectValueAndChange(x interface{}) {
	// v的类型是reflect.Value
	v := reflect.ValueOf(x)
	// 先判断传入的值是不是float32的指针
	// 对于指针，需要调用Elem()函数来判断指针对应的类型
	kind := v.Elem().Kind()
	switch kind {
	case reflect.Float32:
		// 发现指针指向的类型是float32，执行*10
		v.Elem().SetFloat(v.Elem().Float() * 10)
	}
}

func main() {
	var a int32 = 10
	// v的类型：reflect.Value, v的值10
	// 得到的数据类型：int32，值：10
	reflectValue(a)

	// v的类型：reflect.Value, v的值3.2243
	// 得到的数据类型：float32，值：3.2243
	var b float32 = 3.2243
	reflectValue(b)

	var c float32 = 1.25

	////注意需要传入指针，不然会panic
	reflectValueAndChange(&c)
	fmt.Printf("通过反射改变值后，c的值：%v\n", c) // 通过反射改变值后，c的值：12.5
}
