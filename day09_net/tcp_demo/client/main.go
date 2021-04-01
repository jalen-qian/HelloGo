package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

/**
TCP客户端：
1.建立与服务端的连接
2.进行数据收发
3.关闭连接
*/
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	defer conn.Close()
	if err != nil {
		fmt.Printf("建立连接失败，错误：%v\n", err)
		return
	}
	// 不断接收从终端输入的数据，并发送到服务端
	inputReader := bufio.NewReader(os.Stdin)
	for {
		data, rErr := inputReader.ReadString('\n')
		if rErr != nil {
			fmt.Printf("从终端中获取用户输入失败，err:%v\n", rErr)
			continue
		}
		// 对数据进行去空格操作
		data = strings.TrimSpace(data)
		// 如果用户输入"Q"就退出
		if data == "Q" {
			break
		}
		// 发送数据给服务端
		_, wErr := conn.Write([]byte(data))
		if wErr != nil {
			fmt.Printf("发送数据给服务端失败，err:%v\n", wErr)
			continue
		}
		var buf [10]byte
		// 从服务器端接收数据
		n, readErr := conn.Read(buf[:])
		if readErr != nil {
			fmt.Printf("获取服务器端回复失败，err:%v\n", readErr)
			continue
		}
		fmt.Printf("服务器回复：%s\n", string(buf[:n]))
	}
}
