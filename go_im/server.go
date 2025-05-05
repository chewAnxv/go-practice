package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	// 在线用户的列表
	onLineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播的channel
	Message chan string
}

// 构造函数：创建一个Server实例
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		onLineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

// 监听Message广播消息channel的goroutine，一旦有消息就发送给全部在线的User
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message

		// 将msg发送给全部的在线User
		this.mapLock.Lock()
		for _, cli := range this.onLineMap {
			cli.C <- msg

		}
		this.mapLock.Unlock()
	}
}

// 广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg

	this.Message <- sendMsg
}

// 连接处理逻辑，收到客户端连上来就会执行这里
func (this *Server) Handler(conn net.Conn) {

	// fmt.Println("连接建立成功")

	user := NewUser(conn)
	// 用户上线，将用户加入到onlineMap中
	this.mapLock.Lock()
	this.onLineMap[user.Name] = user
	this.mapLock.Unlock()

	// 广播用户上线消息
	this.BroadCast(user, "已上线")

	// 接受客户端发出的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				this.BroadCast(user, "下线")
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}
			// 提取用户的消息（去除'\n'）
			msg := string(buf[:n-1])

			// 将得到的消息进行广播
			this.BroadCast(user, msg)
		}
	}()

	// 当前handler阻塞
	select {}
}

// 启动服务器的接口
func (this *Server) Start() {
	// 1. socket监听（拨号打到交换机）
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))

	if err != nil {
		fmt.Println("net.Listen err：", err)
		return
	}

	defer listener.Close() // 程序结束时关闭监听

	// 启动监听Message的goroutine
	go this.ListenMessager()

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
