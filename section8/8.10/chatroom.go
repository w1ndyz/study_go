package main

import (
	"fmt"
	"net"
	"time"
)

var message = make(chan []byte)

type userInfo struct {
	name    string
	C       chan []byte
	NewUser chan []byte // 用于广播用户进入或退出当前聊天室的消息
}

var onlineUsers = make(map[string]userInfo)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("聊天室-服务已启动")
	fmt.Println("正在监听客户端连接请求")

	// 监听全局channel-message
	go manager()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("地址为[%v]的客户端已连接成功", conn.RemoteAddr())

		go handleConnect(conn)
	}
}

func manager() {
	for {
		select {
		case msg := <-message:
			for _, v := range onlineUsers {
				v.C <- msg
			}
		}
	}
}

func handleConnect(conn net.Conn) {
	defer conn.Close()
	// 用于超时处理
	overTime := make(chan bool)

	// 用于存储用户信息
	buf1 := make([]byte, 4096)
	n, err := conn.Read(buf1)
	if err != nil {
		fmt.Println(err)
		return
	}
	userName := string(buf1[:n])
	perC := make(chan []byte)
	perNewUser := make(chan []byte)
	user := userInfo{name: userName, C: perC, NewUser: perNewUser}
	onlineUsers[conn.RemoteAddr().String()] = user
	fmt.Println("用户[%s]注册成功\n", userName)
	_, _ = conn.Write([]byte("您好," + userName + ", 欢迎您来到聊天室!"))

	// 广播通知,遍历map
	go func() {
		for _, v := range onlineUsers {
			v.NewUser <- []byte("用户[" + userName + "]已加入当前聊天室\n")
		}
	}()

	// 监听每位用户自己的channel
	go func() {
		for {
			select {
			case msg1 := <-user.NewUser:
				_, _ = conn.Write(msg1)
			case msg2 := <-user.C:
				_, _ = conn.Write(msg2)
			}
		}
	}()

	// 循环读取客户端发来的消息
	go func() {
		buf2 := make([]byte, 4096)
		for {
			n, err := conn.Read(buf2)
			// 用于存储当前与服务器通信的客户端上的那个用户名
			thisUser := onlineUsers[conn.RemoteAddr().String()].name
			switch {
			case n == 0:
				fmt.Println(conn.RemoteAddr(), "已断开连接")
				for _, v := range onlineUsers {
					if thisUser != "" {
						v.NewUser <- []byte("用户[" + thisUser + "]已退出当前聊天室\n")
					}
				}
				delete(onlineUsers, conn.RemoteAddr().String())
				return
			case string(buf2[:n]) == "who\n":
				_, _ = conn.Write([]byte("当前在线用户:\n"))
				for _, v := range onlineUsers {
					_, _ = conn.Write([]byte("✅" + v.name + "\n"))
				}
			case len(string(buf2[:n])) > 7 && string(buf2[:n])[:7] == "rename!":
				onlineUsers[conn.RemoteAddr().String()] = userInfo{name: string(buf2[:n-1])[7:], C: perC, NewUser: perNewUser}
				_, _ = conn.Write([]byte("您已成功修改用户名!\n"))
			}

			if err != nil {
				fmt.Println(err)
				return
			}

			var msg []byte
			if buf2[0] != 10 && string(buf2[:n]) != "who\n" {
				if len(string(buf2[:n])) <= 7 || string(buf2[:n])[:7] != "rename!" {
					msg = append([]byte("["+thisUser+"]对大家说:"), buf2[:n]...)
				}
			} else {
				msg = nil
			}

			overTime <- true
			message <- msg
		}
	}()

	for {
		select {
		case <-overTime:
		case <-time.After(time.Second * 60):
			_, _ = conn.Write([]byte("抱歉,由于长时间未发送聊天内容， 您已被踢出系统"))
			thisUser := onlineUsers[conn.RemoteAddr().String()].name
			for _, v := range onlineUsers {
				if thisUser != "" {
					v.NewUser <- []byte("用户[" + thisUser + "]由于长时间未发送消息已被踢出聊天室")
				}
			}
			delete(onlineUsers, conn.RemoteAddr().String())
			return
		}
	}
}
