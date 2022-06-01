package main

import (
	"fmt"
	"net"
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

		fmt.Println("收到数据：", buf[:cnt])

		c.Write(buf[:cnt])
	}

	fmt.Printf("来自 %v 的连接关闭\n", c.RemoteAddr())
}
