package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

// 构造函数：创建一个Server实例
func NewServer(ip string, port int) *Server {
	return &Server{
		Ip:   ip,
		Port: port,
	}
}

// 连接处理逻辑，收到客户端连上来就会执行这里
func (this *Server) Handler(conn net.Conn) {
	fmt.Println("连接建立成功")
}

// 启动服务器
func (this *Server) Start() {
	// 1. socket监听（拨号打到交换机）
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))

	if err != nil {
		fmt.Println("net.Listen err：", err)
		return
	}

	defer listener.Close() // 程序结束时关闭监听

	// 2. 循环接听（接电话）
	for {
		conn, err := listener.Accept() // 阻塞，直到有客户端拨号进来
		if err != nil {
			fmt.Println("listener accept err：", err)
			continue
		}

		// 3. 并发处理（开一个专属协程去跟这个客户端对话
		// 为的是避免一个客服只接听一个电话
		go this.Handler(conn)
	}
}
