package main

import "fmt"

func main() {
	if score := 65; score < 60 {
		fmt.Println("不及格")
	} else if score >= 60 && score < 70 {
		fmt.Println("及格")
	} else if score >= 70 && score < 90 {
		fmt.Println("良好")
	} else {
		fmt.Println("优秀")
	}
}
