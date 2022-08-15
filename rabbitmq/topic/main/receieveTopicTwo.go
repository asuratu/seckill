package main

import (
	"rabbitmq"
	"rabbitmq/topic"
)

func main() {
	r := rabbitmq.NewTopicQueue("topic", "imooc.*.two")
	topic.Consumer(r)
}
