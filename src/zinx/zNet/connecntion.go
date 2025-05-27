package zNet

import (
	"fmt"
	"net"
	ziface "zinxServer/src/zinx/zIface"
)

// 链接模块
type Connection struct {
	//Conn TCP连接
	Conn *net.TCPConn
	//链接ID
	connID uint32
	//当前的状态
	isClosed bool
	//当前链接所绑定的处理业务方法API
	handleAPI ziface.HandleFunc
	//告知当前链接已经退出的channel
	ExitChan chan bool
	//该链接处理的方法
	Router ziface.IRouter
}

// 初始化链接的方法
func NewConnection(conn *net.TCPConn, connID uint32, callback_api ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		connID:    connID,
		isClosed:  false,
		handleAPI: callback_api,
		ExitChan:  make(chan bool, 1),
	}
	return c
}

// 启动链接的都读数据的业务
func (c *Connection) StartReader() {
	fmt.Println("[Zinx] Start Reader Goroutine Running ... ConnID=", c.connID)
	defer fmt.Println("[Zinx] ConnID=", c.connID, " Reader exit!, remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		// 读取数据到buf中，最大读取512字节
		buf := make([]byte, 512)
		n, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("[Zinx] ConnID=", c.connID, " Read data error : ", err)
			continue
		}

		//得到当前conn数据的Request数据
		req := Request{
			conn: c,
			data: buf[:n],
		}

		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
		// 从路由中，找到注册绑定的Conn对应的router调用
	}
}

// 启动链接
func (c *Connection) Start() {
	fmt.Println("[Zinx] Start ConnID=", c.connID)
	//启动从当前链接读数据的业务
	go c.StartReader()
	//TODO: 启动从当前链接写数据的业务
}

// 停止链接
func (c *Connection) Stop() {
	fmt.Println("[Zinx] Stop ConnID=", c.connID)
	//判断当前链接是否关闭
	if c.isClosed {
		return
	}
	//设置当前链接已经关闭
	c.isClosed = true
	//关闭当前链接
	c.Conn.Close()
	//回收资源
	close(c.ExitChan)
}

// 获取当前链接的绑定socket TCP conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// 获取当前的链接ID
func (c *Connection) GetConnID() uint32 {
	return c.connID
}

// 获取远程客户端的TCP状态 IP port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 发送数据，将数据发送到远程的客户端
func (c *Connection) Send(data []byte) error {
	return nil
}
