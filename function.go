package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func Header(conn net.Conn) (to_conn *net.TCPConn) {

	var to_ip []byte
	var to_port []byte
	var to_addr *net.TCPAddr
	header := make([]byte, 25)

	conn.Read(header)
	conn.Write([]byte{0x05, 0x00})

	headerlen, _ := conn.Read(header)
	switch header[3] {
	case 0x01: //ip4
		to_ip = header[4 : 4+net.IPv4len]
	case 0x03: //域名
		ipaddr, _ := net.ResolveIPAddr("ip", string(header[5:headerlen-2])) //如果是域名的话，会返回域名长度，即header[5],函数返回值ipaddr为结构体哦！所以需要下一步
		to_ip = ipaddr.IP
	case 0x04:
		fmt.Println("sorry,目前不支持ipv6哦")
		return
	}
	to_port = header[headerlen-2:]

	conn.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})

	to_addr = &net.TCPAddr{
		IP:   to_ip,
		Port: int(binary.BigEndian.Uint16(to_port)),
	}

	to_conn, _ = net.DialTCP("tcp", nil, to_addr)

	return to_conn
}

func main() {
	l, _ := net.Listen("tcp", ":5678")
	for {
		conn, _ := l.Accept()
		go func() {
			to_conn := Header(conn)
			go func() {
				for {
					data := make([]byte, 4096)
					len, _ := conn.Read(data)
					to_conn.Write(data[:len])
				}
			}()
			for {
				data := make([]byte, 4096)
				len, _ := to_conn.Read(data)
				conn.Write(data[:len])
			}
		}()
	}
}
