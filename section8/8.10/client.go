package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	C    chan string
	Name string
	Addr string
}

var onlineMap map[string]Client

var cmessage = make(chan string)

func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func MakeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + ":" + msg
	return
}

func HandlerConnect(conn net.Conn) {
	defer conn.Close()

	hasData := make(chan bool)

	netAddr := conn.RemoteAddr().String()

	cli := Client{make(chan string), netAddr, netAddr}

	onlineMap[netAddr] = cli

	go WriteMsgToClient(cli, conn)

	cmessage <- MakeMsg(cli, "login")

	isQuit := make(chan bool)

	// 创建一个匿名的go程,专门处理用户发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				isQuit <- true
				fmt.Printf("检测到客户端:%s退出\n", cli.Name)
				return
			}
			if err != nil {
				fmt.Println(err)
				return
			}
			msg := string(buf[:n-1])

			// 提取在线用户列表
			if msg == "who" && len(msg) == 3 {
				conn.Write([]byte("online user list: \n"))
				// 遍历当前map,获取在线用户
				for _, user := range onlineMap {
					userInfo := user.Addr + ":" + user.Name + "\n"
					conn.Write([]byte(userInfo))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				newName := strings.Split(msg, "|")[1]
				cli.Name = newName
				onlineMap[netAddr] = cli
				conn.Write([]byte("rename successful\n"))
			} else {
				cmessage <- MakeMsg(cli, msg)
			}
			hasData <- true
		}
	}()

	// 保证 不退出
	for {
		// 监听channel的数据流动
		select {
		case <-isQuit:
			delete(onlineMap, cli.Addr)
			cmessage <- MakeMsg(cli, "logout")
			return
		case <-hasData:
		case <-time.After(time.Second * 60):
			delete(onlineMap, cli.Addr)
			cmessage <- MakeMsg(cli, "time out leaved")
			return
		}
	}
}
