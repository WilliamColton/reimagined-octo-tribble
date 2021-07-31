package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func main() {
	l, _ := net.Listen("tcp", ":5678")
	for {
		conn, _ := l.Accept()
		go Do(conn)
	}
}

func Do(conn net.Conn) {
	to_conn := To(conn)
	DataCopy(to_conn, conn)
}

func To(conn net.Conn) (to_conn net.Conn) {
	header := make([]byte, 256)

	conn.Read(header)
	conn.Write([]byte{0x05, 0x00})

	io.ReadFull(conn, header[:4])
	atyp := header[3]
	host := ""
	switch atyp {
	case 0x01:
		io.ReadFull(conn, header[:4])
		host = fmt.Sprintf("%d.%d.%d.%d", header[0], header[1], header[2], header[3])
	case 0x03:
		io.ReadFull(conn, header[:1])
		hostlen := int(header[0])
		io.ReadFull(conn, header[:hostlen])
		host = string(header[:hostlen])
	case 0x04:
		fmt.Println("sorry,no IPv6")
		return
	}
	io.ReadFull(conn, header[:2])
	port := binary.BigEndian.Uint16(header[:2])
	addr := fmt.Sprintf("%s:%d", host, port)
	to_conn, _ = net.Dial("tcp", addr)
	conn.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})

	return to_conn
}

func DataCopy(to_conn, conn net.Conn) {
	Copy := func(to_conn, conn net.Conn) {
		defer to_conn.Close()
		defer conn.Close()

		io.Copy(to_conn, conn)
	}

	go Copy(to_conn, conn)
	go Copy(conn, to_conn)
}
