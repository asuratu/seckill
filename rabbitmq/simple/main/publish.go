package main

import (
	"rabbitmq"
	"rabbitmq/simple"
)

func main() {
	r := rabbitmq.NewSimpleQueue("simple")
	// 发送消息
	simple.Publish(r, "hello world 2022")
}
