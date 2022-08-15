package main

import (
	"fmt"
	"rabbitmq"
	"rabbitmq/simple"
	"strconv"
	"time"
)

func main() {
	r1 := rabbitmq.NewRoutingQueue("routing", "routing_key_one")
	r2 := rabbitmq.NewRoutingQueue("routing", "routing_key_two")
	// 发送消息
	for i := 0; i <= 5; i++ {
		simple.Publish(r1, "hello world one "+strconv.Itoa(i))
		simple.Publish(r2, "hello world two "+strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
