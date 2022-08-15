// Package work 工作模式生产者
package work

import (
	"github.com/streadway/amqp"
	"rabbitmq"
)

func Publish(r *rabbitmq.Rabbitmq, message string) {
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
	// 发送消息
	err = r.Channel.Publish(
		r.Exchange, // exchange
		q.Name,     // routing key
		false,      // 如果 mandatory 设置为 true，根据 exchange 和 routing key 没有找到一个合适的 queue，则会返回一个 basic.return 帧给生产者
		false,      // 如果 immediate 设置为 true，当 exchange 发送消息队列后发现没有 consumer 就删除这个 queue
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}
