package main

import (
	"fmt"
	"rabbitmq"
	"rabbitmq/routing"
	"strconv"
	"time"
)

func main() {
	r1 := rabbitmq.NewRoutingQueue("routing", "routing_key_one")
	r2 := rabbitmq.NewRoutingQueue("routing", "routing_key_two")
	// 发送消息
	for i := 0; i <= 5; i++ {
		routing.Publish(r1, "hello world one "+strconv.Itoa(i))
		routing.Publish(r2, "hello world two "+strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
