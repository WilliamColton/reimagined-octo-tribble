package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	header := make([]byte, 25)
	var to_addr []byte
	data := make([]byte, 2048)

	l, _ := net.Listen("tcp", "127.0.0.1:5678")
	for {
		conn, _ := l.Accept()
		conn.Read(header)
		conn.Write([]byte{0x05, 0x00})

		n, _ := conn.Read(header)
		switch header[3] {
		case 0x01:
			to_addr = header[4 : 4+net.IPv4len]
		case 0x03:
			ipAddr, _ := net.ResolveIPAddr("ip", string(header[5:n-2]))
			to_addr = ipAddr.IP
		}
		to_port := header[n-2:]
		Addr := &net.TCPAddr{
			IP:   to_addr,
			Port: int(binary.BigEndian.Uint16(to_port)),
		}
		to, _ := net.DialTCP("tcp", nil, Addr)
		conn.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})

		for {
			conn.Read(data)
			to.Write(data)
			to.Read(data)
			conn.Write(data)
			fmt.Println("ok")
		}

		to.Close()
		conn.Close()
	}
}
