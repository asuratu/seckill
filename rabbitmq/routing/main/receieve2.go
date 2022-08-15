package main

import (
	"rabbitmq"
	"rabbitmq/routing"
)

func main() {
	r := rabbitmq.NewRoutingQueue("routing", "routing_key_two")
	routing.Consumer(r)
}
