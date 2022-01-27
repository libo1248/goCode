package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(fmt.Sprintf("listen error: %s", err))
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("accept error: %s", err)
		}
		go handleConnection(conn)
	}

}

func handleConnection(c net.Conn) {
	//1.conn是否有效
	if c == nil {
		panic("无效的 socket 连接")
	}

	fmt.Printf("来自 %v 的连接建立\n", c.RemoteAddr())

	//2.新建网络数据流存储结构
	buf := make([]byte, 4096)
	//3.循环读取网络数据流
	for {
		//3.1 网络数据流读入 buffer
		cnt, err := c.Read(buf)
		//3.2 数据读尽、读取错误 关闭 socket 连接
		if cnt == 0 || err != nil {
			fmt.Println(cnt, err)
			c.Close()
			break
		}

		//3.3 根据输入流进行逻辑处理
		//buf数据 -> 去两端空格的string
		inStr := strings.TrimSpace(string(buf[0:cnt]))
		//去除 string 内部空格
		cInputs := strings.Split(inStr, " ")
		//获取 客户端输入第一条命令
		fCommand := cInputs[0]

		fmt.Println("客户端传输->" + fCommand)

		switch fCommand {
		case "ping":
			c.Write([]byte("服务器端回复-> pong\n"))
		case "hello":
			c.Write([]byte("服务器端回复-> world\n"))
		default:
			c.Write([]byte("服务器端回复" + fCommand + "\n"))
		}
	}

	fmt.Printf("来自 %v 的连接关闭\n", c.RemoteAddr())
}
