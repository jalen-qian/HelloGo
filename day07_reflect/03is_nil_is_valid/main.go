package main

import (
	"fmt"
	"reflect"
)

/**
1. IsNil()和IsValid()可以用来判断reflect.Value是否是nil以及判断是否持有一个值

2. IsNil()判断Value是否是nil，仅仅只能对引用类型（如通道、函数、接口、映射、指针、切片）来使用否则报panic

3. IsValid()判断Value是否持有一个值，如果Value是Value零值则返回false，在这种情况下，这个Value只能调用
   IsValid()、Kind()、String()这三个方法，其他的都会报panic

用途：IsNil()常被用于判断指针是否为空；IsValid()通常用来判断值类型变量的返回值是否有效
*/

//测试IsNil()方法
func testIsNil(x interface{}) {
	//得到Value
	v := reflect.ValueOf(x)
	//判断IsNil
	if v.IsNil() {
		fmt.Println("是nil")
	} else {
		fmt.Println("不是nil")
	}
}

//测试IsValid
func testIsValid(x interface{}) {
	//获取Value对象
	v := reflect.ValueOf(x)
	if v.IsValid() {
		fmt.Println("有有效的值")
	} else {
		fmt.Println("没有有效的值")
	}
}
func main() {
	var a int32 = 10
	var b *int32 = &a //b是a的指针，int32类型的指针
	testIsNil(b)      //不是nil

	var c *float32
	testIsNil(c) //是nil

	//尝试传入值类型，会报Panic
	//panic: call of reflect.Value.IsNil on int32 Value
	//testIsNil(a)

	//构造一个空切片
	var s1 []int
	testIsNil(s1) //是nil

	//通过make构造一个长度、容量都是0的切片
	var s2 = make([]int, 0, 0)
	testIsNil(s2) //不是nil，说明make分配了内存，切片底层的数组指针不是空指针

	/** 测试IsValid() **/

	aa := ""        //空字符串
	testIsValid(aa) //有有效的值

	var bb *int
	testIsValid(bb) //有有效的值

	//测试nil
	testIsValid(nil) //没有有效的值

	//从map中找不存在的键
	var cc = make(map[string]string)
	cc["娜扎"] = "OK"
	fmt.Println(reflect.ValueOf(cc).MapIndex(reflect.ValueOf("hello")).IsValid()) //false
	fmt.Println(reflect.ValueOf(cc).MapIndex(reflect.ValueOf("娜扎")).IsValid())    //true

	//声明一个匿名结构体（结构体是值类型，这个时候dd是有值的
	var dd struct{ name string }
	//尝试从dd中找name这个属性，并且判断是否存在
	fmt.Println(reflect.ValueOf(dd).FieldByName("name").IsValid()) //true
	fmt.Println(reflect.ValueOf(dd).FieldByName("haha").IsValid()) //false，找不到haha这个属性
	//这个结构体是没有成员方法的，尝试从这个结构体中查找方法
	fmt.Println(reflect.ValueOf(dd).MethodByName("move").IsValid()) //false

	//构造一个invalid reflect.Value
	var v reflect.Value = reflect.ValueOf(dd).FieldByName("hahaha")
	fmt.Printf("%#v\n", v) //<invalid reflect.Value>
	//调用Kind
	v.Kind()
	s := v.String()
	fmt.Println(s)
	//调用非 Kind() String() IsValid()之外的函数，是否报panic
	//panic: reflect: call of reflect.Value.Len on zero Value
	//v.Len()

}
