package main

import (
	"fmt"
	"github.com/tuckersGo/goWeb/web9/cipher"
	"github.com/tuckersGo/goWeb/web9/lzw"
)

type Component interface {
	Operator(string)
}

var sentData string

type SendComponent struct{}

func (self *SendComponent) Operator(data string) {
	sentData = data
}

type ZipComponent struct {
	com Component
}

func (self *ZipComponent) Operator(data string) {
	zipData, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(zipData))
}

type EncryptComponent struct {
	com Component
	key string
}

func (self *EncryptComponent) Operator(data string) {
	encData, err := cipher.Encrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(encData))
}
func main() {
	sender := &EncryptComponent{
		key: "abcde",
		com: &ZipComponent{
			com: &SendComponent{}}}

	sender.Operator("Hello World")

	fmt.Println(sentData)
}
