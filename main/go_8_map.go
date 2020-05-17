package main

import "fmt"

/**
 * Go语言中的Map
 * Hash实现的无序键值对数据结构
 */

func main() {
	/**
	1、定义并使用Map，分为两步
	   1) var map_variable map [map_key_type]map_value_type
	   比如：
	       var students map [int]string
	   这里定义了一个学生map,键是int类型，值是string类型，可以表示学生 学号=>姓名，注意这里声明了一个
	   map，这个map是nil map，不能用来存放键值对
	   2) map_variable := make(map[map_key_type]map_value_type)
	   比如：
	       students := make(map[int]string)
	*/

	//定义一个map,用来存储国家=>首都的关系
	var countryCapitalMap map[string]string

	//可以打印一下，创建的map没有被创建，不能用来存值，运行报错：panic: assignment to entry in nil map
	//countryCapitalMap["China"] = "BeiJing"
	//fmt.Println(countryCapitalMap["china"])//panic: assignment to entry in nil map

	//使用make创建
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["China"] = "BeiJing"
	foreachCountryCapitalMap(countryCapitalMap) //China 的首都是 BeiJing

	//判断HashMap的键是否存在
	findCapital(countryCapitalMap, "China") //China 的首都是 BeiJing
	findCapital(countryCapitalMap, "Japan") //Japan 这个国家的数据不存在

	/**
	如果事先就知道map中的部分数据，则可以直接声明的时候赋值
	*/
	var studentAgeMap = map[string]int{"张三": 23, "李四": 45, "王五": 36}
	/*
		输出：
		张三 的年龄是 23岁
		李四 的年龄是 45岁
		王五 的年龄是 36岁
	*/
	foreachStudentAgeMap(studentAgeMap)
	/**
	Map可以使用delete()函数来删除某个key
	*/
	fmt.Println("删除掉“李四”")
	delete(studentAgeMap, "李四")
	/*
		输出：
		张三 的年龄是 23岁
		王五 的年龄是 36岁
	*/
	foreachStudentAgeMap(studentAgeMap)

}

/**
 * 遍历国家-首都Map
 * 通过range
 */
func foreachCountryCapitalMap(mMap map[string]string) {
	for key, value := range mMap {
		fmt.Printf("%s 的首都是 %s\n", key, value)
	}
}

/**
 * 判断一个Map中是否有某个Key，如果有则打印出首都，没有则判断不存在
 */
func findCapital(mMap map[string]string, country string) {
	capital, ok := mMap[country]
	if ok {
		fmt.Printf("%s 的首都是 %s\n", country, capital) //China 的首都是 BeiJing
	} else {
		fmt.Printf("%s 这个国家的数据不存在\n", country) //Japan 这个国家的数据不存在
	}
}

func foreachStudentAgeMap(studentAgeMap map[string]int) {
	for key, value := range studentAgeMap {
		fmt.Printf("%s 的年龄是 %d岁\n", key, value)
	}
}
