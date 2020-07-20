package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := "hello"
	writer := bufio.NewWriter(os.Stdout)
	_, err := writer.WriteString(s)
	if err != nil {
		fmt.Printf("write to stdout failed, err:%v\n", err)
		return
	}
	writer.Flush()
}
