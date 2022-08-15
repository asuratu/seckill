package main

import (
    "rabbitmq"
    "rabbitmq/simple"
)

func main() {
    r := rabbitmq.NewSimpleQueue("work")
    simple.Consumer(r)
}
