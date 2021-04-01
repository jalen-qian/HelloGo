package main

import (
	"fmt"
	"reflect"
)

/**
Go语言中的反射：

反射是指在程序运行期对程序本身进行访问和修改的能力。程序在编译时，变量被转换为内存地址，变量名不会被编译器写入到可执行部分。在运行程序时，程序无法获取自身的信息。

支持反射的语言可以在程序编译期将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息，这样就可以在程序运行期获取类型的反射信息，并且有能力修改它们。

Go程序在运行期使用reflect包访问程序的反射信息。
*/

/**
1. reflect.TypeOf(x)
reflect.TypeOf可以获取任意变量的【类型对象】
*/

func reflectType(x interface{}) {
	// 不知道传进来的是什么类型？
	// 1.通过类型断言（缺点：只能一个个猜，只适用于事先知道可能传入的是哪几种类型）
	// 2.通过反射 reflect.TypeOf()函数

	t := reflect.TypeOf(x)
	// 类型：*reflect.rtype， 值：int32
	fmt.Printf("t的类型：%T， t的值：%v\n", t, t)

	/*
		可以看到，返回的值是reflect包中的一个结构体(rtype)的指针
		看源码可以看到rtype的定义：

		// rtype is the common implementation of most values.
		// It is embedded in other struct types.
		//
		// rtype must be kept in sync with ../runtime/type.go:/^type._type.
		type rtype struct {
			size       uintptr
			ptrdata    uintptr  // number of bytes in the type that can contain pointers
			hash       uint32   // hash of type; avoids computation in hash tables
			tflag      tflag    // extra type information flags
			align      uint8    // alignment of variable with this type
			fieldAlign uint8    // alignment of struct field with this type
			kind       uint8    // enumeration for C
			alg        *typeAlg // algorithm table
			gcdata     *byte    // garbage collection data
			str        nameOff  // string form
			ptrToThis  typeOff  // type for pointer to this type, may be zero
		}

	*/
}

/**
有时候我们不需要知道变量的具体类型，而是知道变量类型的“分类”，比如是结构体还是函数
是slice还是数组，或者是基本数据类型
这时需要用到 type name和type kind

需要注意的是，Go语言中，数组、map、slice、指针等类型的t.Name()返回的是空
*/
func reflectTypeKind(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Println("类型Name:", t.Name(), " 类型Kind:", t.Kind())
}

type Dog struct {
	name string
}

type Cat struct {
	name string
}

func main() {
	a := int32(12)
	// reflectType(a) //类型：*reflect.rtype， 值：int32

	reflectTypeKind(a) // 类型Name: int32  类型Kind: int32

	b := "hello"
	// reflectType(b) //类型：*reflect.rtype， 值：string

	reflectTypeKind(b) // 类型Name: string  类型Kind: string

	c := [...]string{"hello", "world"}
	// reflectType(c) //类型：*reflect.rtype， 值：[2]string

	reflectTypeKind(c) // 类型Name:   类型Kind: array

	var d Dog
	reflectType(d)     // t的类型：*reflect.rtype， t的值：main.Dog
	reflectTypeKind(d) // 类型Name: Dog  类型Kind: struct

	var c1 Cat
	reflectType(c1)     // t的类型：*reflect.rtype， t的值：main.Cat
	reflectTypeKind(c1) // 类型Name: Cat  类型Kind: struct
}
