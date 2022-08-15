package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

const (
	MQUSER  = "asura"
	MQPWD   = "asura"
	MQHOST  = "localhost"
	MQPORT  = "5672"
	MQVHOST = "imooc"
)

// 格式: amqp://用户名:密码@host:端口/vhost
var mqUrl = fmt.Sprintf("amqp://%s:%s@%s:%s/%s", MQUSER, MQPWD, MQHOST, MQPORT, MQVHOST)

type Rabbitmq struct {
	conn      *amqp.Connection //连接
	Channel   *amqp.Channel    //信道
	QueneName string           //队列名称
	Exchange  string           //交换机
	Key       string           //路由key
	MqUrl     string           //mq地址
}

// NewRabbitMq 创建Rabbitmq结构体实例
func NewRabbitMq(queueName, Exchange, Key string) *Rabbitmq {
	log.Println("mqUrl:", mqUrl)
	return &Rabbitmq{
		QueneName: queueName,
		Exchange:  Exchange,
		Key:       Key,
		MqUrl:     mqUrl,
	}
}

// NewSimpleQueue 创建simple模式队列
func NewSimpleQueue(queueName string) *Rabbitmq {
	q := NewRabbitMq(queueName, "", "")
	q.Connect()
	return q
}

// NewSubscriberQueue 订阅模式创建队列
func NewSubscriberQueue(exchange string) *Rabbitmq {
	q := NewRabbitMq("", exchange, "")
	q.Connect()
	return q
}

// Connect 连接mq
func (r *Rabbitmq) Connect() {
	var err error
	r.conn, err = amqp.Dial(r.MqUrl)
	r.FailOnError(err, "failed to connect to RabbitMQ")
	r.Channel, err = r.conn.Channel()
	r.FailOnError(err, "failed to open a channel")
}

// Destroy 断开连接
func (r *Rabbitmq) Destroy() {
	err := r.Channel.Close()
	r.FailOnError(err, "close channel")
	err = r.conn.Close()
	r.FailOnError(err, "close connection")
}

// FailOnError 错误处理函数
func (r *Rabbitmq) FailOnError(err error, message string) {
	if err != nil {
		// 日志
		log.Printf("%s: %s\n", message, err)
		panic(fmt.Sprintf("%s: %s", message, err))
	}
}
