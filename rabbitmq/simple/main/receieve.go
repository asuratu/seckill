package main

import (
	"rabbitmq"
	"rabbitmq/simple"
)

func main() {
	r := rabbitmq.NewSimpleQueue("simple")
	simple.Consumer(r)
}
