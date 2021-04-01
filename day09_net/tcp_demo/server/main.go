package main

import (
	"fmt"
	"net"
)

/**
3.处理连接
*/
func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		//设置一个缓冲区，从TCP连接中读取到的字节流，会保存到缓冲区中
		//缓冲区是一个byte切片
		/**
		如果缓冲区太小，收到的数据会出现乱码
		收到客户端的数据：hello
		收到客户端的数据：hello 我�
		收到客户端的数据：��小王��
		收到客户端的数据：�
		*/
		var buff [10]byte
		n, err := conn.Read(buff[:])
		if err != nil {
			fmt.Printf("read data failed,err:%v\n", err)
			break
		}
		fmt.Printf("收到客户端的数据：%s\n", string(buff[:n]))
		_, wErr := conn.Write([]byte("数据已收到！"))
		if wErr != nil {
			fmt.Printf("write failed, err:%v\n", wErr)
			break
		}
	}
}

/**
用go语言实现一个tcp连接的服务器端
1.监听端口
2.接收客户端请求并建立连接
3.如果建立成功，则开启一个goroutine处理连接
*/
func main() {
	// 1.监听本地的20000端口
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("listen failed,err:%v\n", err)
		return
	}

	// 无限循环，不断的去监听客户端是否发起连接
	for {
		conn, err := listen.Accept() // 接收连接
		if err != nil {
			fmt.Printf("accept connection failed, err:%v\n", err)
		}
		// 2.成功建立连接，创建一个goroutine去处理
		go handleConn(conn)
	}
}
