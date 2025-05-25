package zNet

import (
	"errors"
	"fmt"
	"net"
	ziface "zinxServer/src/zinx/zIface"
)

type Server struct {
	//服务器名称
	Name string
	//服务器绑定的IP版本
	IPVersion string
	//服务器IP地址
	IP string
	//服务器端口
	Port int
	//最大连接数
	MaxConn int
	//当前连接数
	CurConn int
}

// NewServer 创建一个新的Server实例
func NewServer(name string) ziface.IServer {
	s := new(Server)
	s.Name = name
	s.IPVersion = "tcp4"
	s.IP = "0.0.0.0"
	s.Port = 8080
	return s
}

// 定义当前客户端所绑定的API回调函数，
func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	// 直接回写数据到客户端
	_, err := conn.Write(data[:cnt])
	if err != nil {
		fmt.Println("CallBackToClient Write error:", err)
		return errors.New("CallBackToClient Write error")
	}
	return nil
}

// 启动服务器
func (s *Server) Start() {
	fmt.Printf("[Zinx] Server Name: %s, Listening on %s:%d\n", s.Name, s.IP, s.Port)

	// 1. 获取一个TCP Addr
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("ResolveTCPAddr err:", err)
		return
	}
	go func() {
		// 2. 启动监听
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("ListenTCP err:", err)
			return
		}
		fmt.Println("Start Zinx server success, ", s.Name, " is listening...")
		var connID uint32
		connID = 0
		// 3. 循环接收客户端连接（读写）
		for {
			// 如果客户端连接过来，阻塞会返回
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err:", err)
				continue
			}
			// 4. 创建一个新的连接对象
			Connection := NewConnection(conn, connID, CallBackToClient)
			//启动业务
			go Connection.Start()
			connID++
		}
	}()
}

// 停止服务器
func (s *Server) Stop() {
	//TODO: 做一些停止服务器之前的业务
}

// 运行服务器
func (s *Server) Server() {
	// 1. 启动服务器
	s.Start()

	//TODO: 做一些启动服务器之后额外的业务

	// 阻塞主线程
	select {}
}
