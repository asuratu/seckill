// Package simple 简单模式消费者
package simple

import (
	"fmt"
	"rabbitmq"
)

func Consumer(r *rabbitmq.Rabbitmq) {
	// 申请队列，如果队列不存在，则自动创建
	q, err := r.Channel.QueueDeclare(
		r.QueneName, // name
		false,       // 是否持久化
		false,       // 是否自动删除
		false,       // 是否具有排他性，只有当前用户可以访问
		false,       // 是否阻塞
		nil,         // 额外属性
	)
	r.FailOnError(err, "Failed to declare a queue")
	// 消费消息
	messages, err := r.Channel.Consume(
		q.Name, // 队列名称
		"",     // 用来区分多个消费者
		true,   // 是否自动应答，false 的话需要手动应答（回调函数），true 的话会自动应答
		false,  // 是否具有排他性，只有当前用户可以访问
		false,  // 如果为 true，表示不能将同一个 connection 的消费者放在不同的 goroutine 中
		false,  // 是否阻塞，
		nil,    // 额外属性
	)
	r.FailOnError(err, "Failed to register a consumer")
	forever := make(chan bool)
	go func() {
		for d := range messages {
			// 消费消息
			fmt.Printf("Received a message: %s\n", d.Body)
		}
	}()
	fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
