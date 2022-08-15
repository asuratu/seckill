package main

import (
	"fmt"
	"rabbitmq"
	"rabbitmq/simple"
	"strconv"
	"time"
)

func main() {
	r := rabbitmq.NewSimpleQueue("work")
	// 发送消息
	for i := 0; i <= 5; i++ {
		simple.Publish(r, "hello world "+strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
