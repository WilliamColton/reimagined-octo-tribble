package main

import (
	"encoding/binary"
	"fmt"
	"log"

	"github.com/gogf/gf/v2/crypto/gaes"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtcp"
)

func Handle(c *gtcp.Conn) {
	key := make([]byte, 16)
	for i := range key {
		key[i] = byte(i)
	}

	// 错误处理函数，简化错误处理代码
	handleError := func(err error, msg string) bool {
		if err != nil {
			log.Println(msg, err)
			c.Close() // 出现错误时关闭客户端连接
			return true
		}
		return false
	}

	data, err := c.Recv(-1)
	if handleError(err, "接收数据失败:") {
		return
	}

	data, err = gaes.Encrypt([]byte{0x05, 0x00}, key)
	if handleError(err, "数据加密失败:") {
		return
	}
	if handleError(c.Send(data), "发送数据失败:") {
		return
	}

	data, err = c.Recv(-1)
	if handleError(err, "接收数据失败:") {
		return
	}
	data, err = gaes.Decrypt(data, key)
	if handleError(err, "数据解密失败:") {
		return
	}

	datalen := len(data)
	addr := data[5 : datalen-2]
	port := data[datalen-2:]
	address := fmt.Sprintf("%v:%v", string(addr), binary.BigEndian.Uint16(port))
	fmt.Println("解析地址:", address)

	data, err = gaes.Encrypt([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, key)
	if handleError(err, "数据加密失败:") {
		return
	}
	if handleError(c.Send(data), "发送数据失败:") {
		return
	}

	to, err := gtcp.NewConn(address)
	if handleError(err, "连接到远程地址失败:") {
		return
	}
	defer to.Close()

	go func() {
		for {
			buf, err := c.Recv(-1)
			if err != nil {
				log.Println("从客户端接收失败:", err)
				return
			}
			buf, err = gaes.Decrypt(buf, key)
			if err != nil {
				log.Println("解密失败:", err)
				return
			}
			if err := to.Send(buf); err != nil {
				log.Println("发送到远程服务器失败:", err)
				return
			}
		}
	}()

	for {
		buf, err := to.Recv(-1)
		if err != nil {
			log.Println("从远程服务器接收失败:", err)
			return
		}
		buf, err = gaes.Encrypt(buf, key)
		if err != nil {
			log.Println("加密失败:", err)
			return
		}
		if err := c.Send(buf); err != nil {
			log.Println("发送到客户端失败:", err)
			return
		}
	}
}

func main() {
	Server := g.TCPServer("Server")
	//注意事项：在运行时阶段，每一次通过g模块获取单例对象时都会有内部全局锁机制来保证操作和数据的并发安全性，原理性上来讲在并发量大的场景下会存在锁竞争的情况，但绝大部分的业务场景下开发者均不需要太在意锁竞争带来的性能损耗。此外，开发者也可以通过将获取到的单例对象保存到特定的模块下的内部变量重复使用，以此避免运行时锁竞争情况。
	Server.SetAddress("127.0.0.1:8900")
	Server.SetHandler(Handle)
	if err := Server.Run(); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
