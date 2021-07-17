package main

import (
	"fmt"
	"io"
	"net"
	"runtime"
	"strconv"
)

func main() {
	fmt.Println("start")
	l, _ := net.Listen("tcp", "127.0.0.1:5678")
	defer l.Close()
	runtime.GOMAXPROCS(runtime.NumCPU())
	for {
		func() {
			var n int
			var host string
			var port string

			c, _ := l.Accept()
			defer c.Close()

			var header = make([]byte, 25)

			c.Read(header)
			c.Write([]byte{0x05, 0x00})

			n, _ = c.Read(header)
			atyp := header[3]
			switch atyp {
			case 0x01:
				host = net.IPv4(header[4], header[5], header[6], header[7]).String()
			case 0x03:
				host = string(header[5 : n-2]) //atpy=0x03时,header[4]为域名长度
			case 0x04:
				return
			}
			port = strconv.Itoa(int(header[n-2])<<8 | int(header[n-1]))

			to, _ := net.Dial("tcp", net.JoinHostPort(host, port))
			defer to.Close()
			c.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})

			io.Copy(c, to)
			io.Copy(to, c)
			fmt.Print(runtime.NumGoroutine())
		}()
	}
}
