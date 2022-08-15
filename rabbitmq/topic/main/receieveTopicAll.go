package main

import (
	"rabbitmq"
	"rabbitmq/topic"
)

func main() {
	r := rabbitmq.NewTopicQueue("topic", "#")
	topic.Consumer(r)
}
