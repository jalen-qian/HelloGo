package main

import (
	"fmt"
	"reflect"
)

/**
结构体反射
*/
type Student struct {
	Name  string `json:"name" ini:"s_name"`
	Age   int    `json:"age" ini:"s_age"`
	Score int    `json:"score" ini:"s_score"`
}

func (s Student) Study() (m string) {
	msg := "我在学习课程"
	fmt.Println(msg)
	m = msg
	return
}

func (s Student) Sleep() string {
	msg := "夜深了，该睡觉了"
	fmt.Println(msg)
	return msg
}

/** 学生打游戏，需要传入游戏名称 **/
func (s Student) Game(game string) {
	fmt.Println("我正在打", game, "游戏")
}

/**
通过反射获取结构体的字段信息
*/
func reflectField(x interface{}) {
	typeObj := reflect.TypeOf(x)
	// 遍历每个字段，得到字段的信息
	for i := 0; i < typeObj.NumField(); i++ {
		fieldObj := typeObj.Field(i) // 根据索引得到属性
		/*
			name: Name pkgPath: type:string tag:json:"name" ini:"s_name"
			name: Age pkgPath: type:int tag:json:"age" ini:"s_age"
			name: Score pkgPath: type:int tag:json:"score" ini:"s_score"
		*/
		fmt.Printf("name: %v pkgPath:%v type:%v tag:%v\n", fieldObj.Name, fieldObj.PkgPath, fieldObj.Type, fieldObj.Tag)
		// 可以看到tag有两个，如果需要得到其中具体的一个tag
		fmt.Println("jsonTag:", fieldObj.Tag.Get("json"))
	}

	// 还可以根据字段名称获取字段信息
	fieldObj, ok := typeObj.FieldByName("Score")
	if ok {
		fmt.Println("根据名称获取字段信息：")
		fmt.Printf("name: %v pkgPath:%v type:%v tag:%v\n", fieldObj.Name, fieldObj.PkgPath, fieldObj.Type, fieldObj.Tag)
	}
}

func reflectMethod(x interface{}) {
	t := reflect.TypeOf(x)
	// 获取方法数量
	fmt.Printf("num of method:%d\n", t.NumMethod()) // num of method:2
	// 因为接下来需要获取方法的一些信息，并调用方法，所以用ValueOf
	v := reflect.ValueOf(x)
	//ValueOf()同样可以获得方法数量，这个是因为Value结构体嵌套了 rtype结构体，rtype实现了Type接口
	//for i := 0; i < v.NumMethod(); i++ {
	//	//通过索引获取method
	//	//注意这里获取方法名称是通过TypeOf获取，t.Method(i)返回的是Method结构体
	//	//而v.Method(i)返回的是Value结构体
	//	fmt.Printf("methodName:%v ", t.Method(i).Name)
	//	fmt.Printf("methodType:%v\n", t.Method(i).Type)
	//	fmt.Printf("根据v获取methodType:%v\n", v.Method(i).Type())
	//
	//	//调用函数,注意一定要传入一个reflect.Value类型的切片
	//	var args []reflect.Value
	//	v.Method(i).Call(args)
	//}

	// 反射的方法如果需要传参数，如何传呢？
	args := make([]reflect.Value, 1)
	// 构造参数
	args[0] = reflect.ValueOf("绝地求生")
	v.MethodByName("Game").Call(args)
}

func main() {
	stu1 := Student{
		Name:  "小王子",
		Age:   18,
		Score: 96,
	}
	// 反射字段
	// reflectField(stu1)

	// 反射方法
	reflectMethod(stu1)
}
