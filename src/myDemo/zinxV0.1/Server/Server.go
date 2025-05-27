package main

import (
	"fmt"
	ziface "zinxServer/src/zinx/zIface"
	"zinxServer/src/zinx/zNet"
)

type PingRouter struct {
	zNet.BaseRouter
}

func (pr *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping..."))
	if err != nil {
		fmt.Println("call before handle error:", err)
	}
}

func (pr *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping .. ping.. ping.."))
	if err != nil {
		fmt.Println("call Handle error:", err)
	}
}

func (pr *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping..."))
	if err != nil {
		fmt.Println("call after handle error:", err)
	}
}

func main() {
	s := zNet.NewServer("[Zinx V0.1]")
	//添加自定义的Router
	pr := new(PingRouter)
	s.AddRouter(pr)
	//启动Server
	s.Server()
}
