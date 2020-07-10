package tools

import "fmt"

func printIntSlice(s []int, name string) {
	fmt.Printf("len(%s)=%d, cap(%s)=%d,%s=%v,地址：%p\n", name, len(s), name, cap(s), name, s, s)
}
