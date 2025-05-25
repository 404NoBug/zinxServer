package main

import (
	"fmt"
	"net"
	"time"
)

// 模拟客户端
func main() {
	fmt.Println("Client start...")
	time.Sleep(1 * time.Second) // 模拟客户端启动延时
	// 1. 创建一个TCP连接
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	defer conn.Close()
	// 2. 发送数据
	for {
		_, err = conn.Write([]byte("Hello Zinx V0.1.."))
		if err != nil {
			fmt.Println("Write err:", err)
			return
		}

		// 3. 接收数据
		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Println("Read err:", err)
			return
		}
		fmt.Println("Client recv:", string(buff[:n]))

		// cup阻塞
		time.Sleep(1 * time.Second)
	}
}
