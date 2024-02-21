package main

import (
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

	to, err := gtcp.NewConn(":8900")
	if err != nil {
		log.Println("连接到远程服务器失败:", err)
		c.Close() // 确保关闭客户端连接
		return
	}
	defer to.Close()
	defer c.Close()

	go func() {
		for {
			buf, err := c.Recv(-1)
			if err != nil {
				log.Println("从客户端接收失败:", err)
				break // 出错时退出循环
			}
			encryptedBuf, err := gaes.Encrypt(buf, key)
			if err != nil {
				log.Println("加密失败:", err)
				break // 出错时退出循环
			}
			if err := to.Send(encryptedBuf); err != nil {
				log.Println("发送到服务器失败:", err)
				break // 出错时退出循环
			}
		}
		c.Close() // 出错时确保关闭连接
		to.Close()
	}()

	for {
		buf, err := to.Recv(-1)
		if err != nil {
			log.Println("从服务器接收失败:", err)
			break // 出错时退出循环
		}
		decryptedBuf, err := gaes.Decrypt(buf, key)
		if err != nil {
			log.Println("解密失败:", err)
			break // 出错时退出循环
		}
		if err := c.Send(decryptedBuf); err != nil {
			log.Println("发送到客户端失败:", err)
			break // 出错时退出循环
		}
	}
}

func main() {
	Local := g.TCPServer("Local")
	//注意事项：在运行时阶段，每一次通过g模块获取单例对象时都会有内部全局锁机制来保证操作和数据的并发安全性，原理性上来讲在并发量大的场景下会存在锁竞争的情况，但绝大部分的业务场景下开发者均不需要太在意锁竞争带来的性能损耗。此外，开发者也可以通过将获取到的单例对象保存到特定的模块下的内部变量重复使用，以此避免运行时锁竞争情况。
	Local.SetAddress("127.0.0.1:5678")
	Local.SetHandler(Handle)
	if err := Local.Run(); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
