package main

import "fmt"

func printSlice(s []string, name string) {
	fmt.Printf("len(%s)=%d, cap(%s)=%d,%s=%v,地址：%p\n", name, len(s), name, cap(s), name, s, s)
}

func main() {
	s1 := []string{"北京", "上海", "深圳"}
	// s1[3] = "广州" //index out of range
	printSlice(s1, "s1") // len(s1)=3, cap(s1)=3,s1=[北京 上海 深圳],地址：0xc000058360
	s1 = append(s1, "广州")
	// 长度和容量都变了，说明append之后，底层数组拷贝到了新的地址
	// 注意，只有当底层数组容量不够时，才会整体拷贝分配新地址（搬家）
	printSlice(s1, "s1") // len(s1)=4, cap(s1)=6,s1=[北京 上海 深圳 广州],地址：0xc000044060

	// 由于底层数组容量扩充到了6，所以再拼接两个元素，不会重新分配地址
	s1 = append(s1, "杭州", "重庆")
	// len(s1)=6, cap(s1)=6,s1=[北京 上海 深圳 广州 杭州 重庆],地址：0xc000044060
	printSlice(s1, "s1")

	// 定义一个s2，并将s2中的所有元素通过append函数拼接到s1的后面
	s2 := []string{"武汉", "南京", "拉萨"}
	//"..."语法：s2是一个切片,append第二个参数需要传字符串类型，...将切片拆开为3个字符串
	s1 = append(s1, s2...)

	// len(s1)=9, cap(s1)=12,s1=[北京 上海 深圳 广州 杭州 重庆 武汉 南京 拉萨],地址：0xc0000960c0
	printSlice(s1, "s1")
}
