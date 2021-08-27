package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	Getarray()
}

func Getarray() (array []byte) {
	array = make([]byte, 256)
	for i := 0; i < 256; i++ {
		array[i] = byte(rand.Intn(256))
		for key, _ := range array {
			if array[i] == array[key] {
				return Getarray() //一定要加return,不然运行一会之后直接提示stack overflow(栈溢出) QAQ~.额，经过测试，如果这样写加了return也会溢出QAQ。
			}
		}
	}
	fmt.Println(array)
	fmt.Println(len(array))

	return array
}

func Key(array []byte) (enkey []byte, dekey []byte) {
	return
}

func Endata(enkey []byte, data []byte) (endata []byte) {
	return
}

func Dedata(dekey []byte, endata []byte) (dedata []byte) {
	return
}
