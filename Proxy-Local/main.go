package main

import (
	"fmt"

	"github.com/gogf/gf/v2/crypto/gaes"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtcp"
)

func Handle(c *gtcp.Conn) {
	/*
		data, _ := c.Recv(-1)
		fmt.Println(data)
		c.Send([]byte{0x05, 0x00})

		data, _ = c.Recv(-1)

		datalen := len(data)
		addr := data[5 : datalen-2]
		port := data[datalen-2:]
		address := fmt.Sprintf("%v:%v", gconv.String(addr), gconv.String(binary.BigEndian.Uint16(port)))
		fmt.Println(address)

		c.Send([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})

		to, _ := gtcp.NewConn(address)
	*/
	to, err := gtcp.NewConn(":8900")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connect Success")
	}
	defer c.Close()
	defer to.Close()

	go func() {
		for {
			buf, _ := c.Recv(-1)
			buf, err := gaes.Encrypt(gaes.PKCS7Padding(buf, 16), []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 12, 3, 4, 5, 65})
			if err != nil {
				fmt.Println(err)
			}
			to.Send(buf)
		}
	}()
	for {
		buf, _ := to.Recv(-1)
		buf, err := gaes.Decrypt(buf, []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 12, 3, 4, 5, 65})
		if err != nil {
			fmt.Println(err)
		}
		buf, _ = gaes.PKCS7UnPadding(buf, 16)
		c.Send(buf)
	}
}

func main() {
	Local := g.TCPServer("Local")
	//注意事项：在运行时阶段，每一次通过g模块获取单例对象时都会有内部全局锁机制来保证操作和数据的并发安全性，原理性上来讲在并发量大的场景下会存在锁竞争的情况，但绝大部分的业务场景下开发者均不需要太在意锁竞争带来的性能损耗。此外，开发者也可以通过将获取到的单例对象保存到特定的模块下的内部变量重复使用，以此避免运行时锁竞争情况。
	Local.SetAddress("127.0.0.1:5678")
	Local.SetHandler(Handle)
	Local.Run()
}
