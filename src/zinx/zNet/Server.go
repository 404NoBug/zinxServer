package zNet

import (
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
	//给当前的server添加一个router，server注册的链接对应的处理业务
	Router ziface.IRouter

	// //最大连接数
	// MaxConn int
	// //当前连接数
	// CurConn int
}

// 路由功能：给当前的服务器注册一个路由方法，供客户端的链接处理使用
func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("[Zinx] Add Router to Server:", s.Name)
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
		var connID uint32 = 0
		// 3. 循环接收客户端连接（读写）
		for {
			// 如果客户端连接过来，阻塞会返回
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err:", err)
				continue
			}
			// 4. 创建一个新的连接对象
			Connection := NewConnection(conn, connID, s.Router)

			connID++
			//启动业务
			go Connection.Start()
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

// NewServer 创建一个新的Server实例
func NewServer(name string) ziface.IServer {
	s := new(Server)
	s.Name = name
	s.IPVersion = "tcp4"
	s.IP = "0.0.0.0"
	s.Port = 8080
	s.Router = nil // 初始化Router为nil
	return s
}
