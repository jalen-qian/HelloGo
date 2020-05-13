package main
import ("fmt")

func main(){
	//定义切片（不固定长度的数组）
	var a [] int = []int{1,3,4}
	fmt.Println(a)
	//使用make函数创建切片
	var b = make([]int,5)
	b[0] = 2
	fmt.Println(b)
}