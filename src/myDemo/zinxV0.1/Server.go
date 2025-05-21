package main

import (
	"zinx/zNet"
)

func main() {
	s := zNet.NewServer("[Zinx V0.1]")
	s.Server()
}