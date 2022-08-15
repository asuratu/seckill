// Package topic 话题模式消费者
package topic

import (
	"log"
	"rabbitmq"
)

// Consumer 路由模式消费者
func Consumer(r *rabbitmq.Rabbitmq) {
	// 尝试创建交换机
	err := r.Channel.ExchangeDeclare(
		r.Exchange, // 交换机名称
		"topic",    // 交换机类型，topic 是话题模式，所有消息都会到达所有的 queue
		true,       // 是否持久化
		false,      // 是否自动删除，如果设置为 true，则 queue 和 exchange 会在服务器重启后自动删除
		false,      // 是否具有排他性，只有当前用户可以访问
		false,      // 是否阻塞
		nil,        // 额外属性
	)
	r.FailOnError(err, "Failed to declare an exchange")
	// 尝试创建队列
	q, err := r.Channel.QueueDeclare(
		"",    // 队列名称，随机生成
		false, // 是否持久化
		false, // 是否自动删除，如果设置为 true，则 queue 和 exchange 会在服务器重启后自动删除
		true,  // 是否具有排他性，只有当前用户可以访问
		false, // 是否阻塞
		nil,   // 额外属性
	)
	r.FailOnError(err, "Failed to declare a queue")
	// 绑定队列到交换机
	err = r.Channel.QueueBind(
		q.Name,     // 队列名称
		r.Key,      // 话题模式的 key 规则：* 用于匹配一个单词，# 用于匹配多个单词（可以是零个），例如：imooc.* 表示匹配 imooc.hello，imooc.# 表示匹配 imooc.hello.world
		r.Exchange, // exchange
		false,      // 是否排他性
		nil,        // 额外属性
	)
	r.FailOnError(err, "Failed to bind a queue")
	// 消费消息
	msgs, err := r.Channel.Consume(
		q.Name, // 队列名称
		"",     // consumer 名称，随机生成
		true,   // 是否自动应答
		false,  // 是否具有排他性，只有当前用户可以访问
		false,  // 是否阻塞
		false,  // 如果设置为 true，则 queue 和 exchange 会在服务器重启后自动删除
		nil,    // 额外属性
	)
	r.FailOnError(err, "Failed to register a consumer")
	// 消费消息
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Routing received a message from %s: %s\n", r.Key, d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C\n")
	<-forever
}
