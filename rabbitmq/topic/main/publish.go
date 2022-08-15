package main

import (
	"fmt"
	"rabbitmq"
	"rabbitmq/topic"
	"strconv"
	"time"
)

func main() {
	r1 := rabbitmq.NewTopicQueue("topic", "all.topic")
	r2 := rabbitmq.NewTopicQueue("topic", "imooc.topic.two")
	// 发送消息
	for i := 0; i <= 5; i++ {
		topic.Publish(r1, "hello world all "+strconv.Itoa(i))
		topic.Publish(r2, "hello world two "+strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
