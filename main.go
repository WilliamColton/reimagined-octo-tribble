package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	data := []byte{0, 1, 2}
	data = endata(data)
	fmt.Println(data, dedata(data))
}

func endata(data []byte) (_data []byte) {
	data_len := len(data)
	_data = make([]byte, data_len)

	for i, v := range data {
		_data[data_len-i-1] = v
		//fmt.Println(v, data_len-i-1)
		//fmt.Println(_data, _data[data_len-1])
	}
	//data=_data
	//fmt.Println(data, _data)//此处说明函数内的值并未影响到函数外，离谱。。。 只好加个返回值去处理咯。
	return _data
	//fine!
}

func dedata(data []byte) (_data []byte) {
	data_len := len(data)
	_data = make([]byte, data_len)

	for i, v := range data {
		_data[data_len-i-1] = v
	}

	return _data
	//fine!
}

func decoderead(c net.Conn, data []byte) (len int, err error, _data []byte) {
	len, err = c.Read(data)
	if err != nil {
		return
	}
	_data = dedata(data[:len])

	return len, err, _data
}

func encodewrite(c net.Conn, data []byte) (len int, err error) {
	_data := endata(data)

	return c.Write(_data)
}

func encopy(s net.Conn, c net.Conn) error {
	data := make([]byte, 512)
	for {
		readlen, readerr := s.Read(data)
		if readerr != nil {
			if readerr != io.EOF {
				return readerr
			} else {
				return nil
			}
		}
		if readlen > 0 {
			writelen, writeerr := encodewrite(c, data[:readlen])
			if writeerr != nil {
				return writeerr
			}
			if readlen != writelen {
				return io.ErrShortWrite
			}
		}
	}
}

func decopy(s net.Conn, c net.Conn) error {
	data := make([]byte, 512)
	for {
		readlen, readerr, data := decoderead(s, data)
		if readerr != nil {
			if readerr != io.EOF {
				return readerr
			} else {
				return nil
			}
		}
		if readlen > 0 {
			writelen, writeerr := c.Write(data[:readlen])
			if writeerr != nil {
				return writeerr
			}
			if readlen != writelen {
				return io.ErrShortWrite
			}
		}
	}
}
