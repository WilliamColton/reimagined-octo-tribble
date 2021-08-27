//加解密部分终于在2021/8/27日下午16:45完成
//先把函数封装好，避免后期修改困难QAQ
//注意标点符号,全部用英文标点!!!-----------------------------------------------------
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	data := []byte{0, 2, 5, 8, 77}
	rand.Seed(time.Now().UnixNano())
	a := Getarray()
	en, de := Key(a)
	endata := Endata(en, data)
	fmt.Println(Endata(en, data))
	fmt.Println(Dedata(de, endata))
}

//	以下所有数字都不要更改!!!!!!----------------------------------------------------

func Getarray() (array []byte) {
	array = make([]byte, 256)
	int_array := rand.Perm(256)

	for key, value := range int_array {
		array[key] = byte(value)
		if key == value { //	保证索引和值不会出现重复的情况，但是也可能像英格玛密码机一样变成弱点
			return Getarray() //	此处加不加return无所谓，但是加了更美观
		}
	}

	return array
}

func Key(array []byte) (enkey []byte, dekey []byte) {
	enkey = make([]byte, 256)
	dekey = make([]byte, 256)

	for key, value := range array {
		dekey[value] = byte(key)
	}

	enkey = array

	return enkey, dekey
}

//	以上所有数字都不要更改!!!-------------------------------------------------------

func Endata(enkey []byte, data []byte) (endata []byte) {
	datalen := len(data)
	endata = make([]byte, datalen)

	for key, value := range data {
		endata[key] = enkey[value]
	}

	return endata
}

func Dedata(dekey []byte, endata []byte) (dedata []byte) {
	endatalen := len(endata)
	dedata = make([]byte, endatalen)

	for key, value := range endata {
		dedata[key] = dekey[value]
	}

	return dedata
}
