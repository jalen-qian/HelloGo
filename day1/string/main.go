package main

import "fmt"

func main() {
	dir := "C:\\Windows\\System32\\drivers\\etc"
	fmt.Println(dir) //C:\Windows\System32\drivers\etc

	/*

			第一行
			第二行
		第三行
				哈哈哈
	*/
	var multiLine = `
		第一行
		第二行
	第三行
			哈哈哈
    
    `
	fmt.Println(multiLine)
}
