package main

import (
	"rabbitmq"
	"rabbitmq/subscriber"
)

func main() {
	r := rabbitmq.NewSubscriberQueue("subscriber")
	subscriber.Consumer(r)
}
