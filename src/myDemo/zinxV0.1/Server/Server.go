package main

import (
	"zinxServer/src/zinx/zNet"
)

func main() {
	s := zNet.NewServer("[Zinx V0.1]")
	s.Server()
}
