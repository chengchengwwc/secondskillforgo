package rabbitmq

import (
	"fmt"
	"log"
	"github.com/streadway/amqp"
)


// url格式 amqp://账号:密码@rabbitmq服务器地址:端口号/vhost
const MQURL = "amqp://guest:guest@127.0.0.1:5672/imooc"

type RabbitMQ struct{
	conn *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机
	Exchange string
	//key
	Key string
	//连接信息
	Mqurl string
}

func NewRabbitMq(queueName string,exchange string,key string) *RabbitMQ{
	return &RabbitMQ{QueueName: queueName,Exchange:exchange,Key:key,Mqurl: MQURL}
}

func ( r *RabbitMQ) Destory(){
	r.channel.Close()
	r.conn.Close()
}

func (r *RabbitMQ) failOnErr(err error,message string){
	if err != nil{
		log.Fatalf("%s:%s",message,err)
		panic(fmt.Sprintf("%s:%s",message,err))
	}
}

// simple work
func NewRabbitMQSimple(queueName string) *RabbitMQ{
	rabbitmq := NewRabbitMq(queueName,"","")
	var err error
	rabbitmq.conn,err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err,"链接生成失败")
	rabbitmq.channel,err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err,"获取channel失败")
	return rabbitmq
}
// 简单模式下生产代码
func (r *RabbitMQ) PublishSimple(message string){
	// 1. 申请队列 如果队列不存在，则自动创建，如果存在则跳过创建
	// 保证队列存在，消息能发送到队列中
	_,err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否为自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞
		false,
		//额外属性
		nil,
	)
	if err != nil{
		fmt.Println(err)
		return
	}
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		//如果为true会根据exchange类型和routekey规则，如果无法找到符合条件的规则，则会把瞎消息发送给其他消费者
		false,
		//如果为true,当exchange 发送消息到消息队列后，如果没有绑定消费者，则会将消息返回给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:[]byte(message),
		},
	)
}

func (r *RabbitMQ) ConsumeSimple(){
	// 1. 申请队列 如果队列不存在，则自动创建，如果存在则跳过创建
	// 保证队列存在，消息能发送到队列中
	_,err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否为自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞
		false,
		//额外属性
		nil,
	)
	if err != nil{
		fmt.Println(err)
		return
	}
	//接受消息
	msgs,err := r.channel.Consume(
		r.QueueName,
		//用来区分多个消费者
		"",
		//是否自动应答
		true,
		//是否具有排他性
		false,
		//如果设置为true, 表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,
		//队列消费是否为阻塞
		false,
		nil,
	)
	if err != nil{
		fmt.Println(err)
		return
	}
	//启用携程来处理函数
	forver := make(chan bool)
	go func(){
		for d:= range msgs{
			//实现我们要处理的逻辑函数
			log.Printf("resevie message %s",d.Body)
		}
	}()
	log.Printf("[*] Wating for message,to exit press ctrl +c")
	<- forver
}


//订阅模式生产
func (r *RabbitMQ) PublishPub(message string){
	//1. 尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		//交换机类型
		"fanout",
		//是否持久化
		true,
		//是否删除
		false,
		//
		false,
		false,
		nil,
	)
	if err != nil{
		return
	}
	//2. 发送消息
	err = r.channel.Publish(
		r.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(message),
		},
	)
}
//订阅模式下消费者
func(r *RabbitMQ) RecieveSub(){
	//1. 尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		//交换机类型
		"fanout",
		//是否持久化
		true,
		//是否删除
		false,
		//
		false,
		false,
		nil,
	)
	if err != nil{
		return
	}
	//创建队列
	q,err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	//绑定队列到exchange上
	err = r.channel.QueueBind(
		q.Name,
		//在pub/sub模式下，这里的key要为空
		"",
		r.Exchange,
		false,
		nil,
	)

	message, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil{
		return
	}

	//启用携程来处理函数
	forver := make(chan bool)
	go func(){
		for d:= range message{
			//实现我们要处理的逻辑函数
			log.Printf("resevie message %s",d.Body)
		}
	}()
	log.Printf("[*] Wating for message,to exit press ctrl +c")
	<- forver
}

func (r *RabbitMQ) PublishRouting(message string){
	//1. 尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		//交换机类型
		"direct",
		//是否持久化
		true,
		//是否删除
		false,
		//
		false,
		false,
		nil,
	)
	if err != nil{
		return
	}
	//2. 发送消息
	err = r.channel.Publish(
		r.Exchange,
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(message),
		},
	)

}

func (r *RabbitMQ) RecieveRouting(){
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		//交换机类型
		"direct",
		//是否持久化
		true,
		//是否删除
		false,
		//
		false,
		false,
		nil,
	)
	if err != nil{
		return
	}
	q,err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	//绑定队列到exchange上
	err = r.channel.QueueBind(
		q.Name,
		//在routing模式下，要设置key
		r.Key,
		r.Exchange,
		false,
		nil,
	)
	if err != nil{
		return 
	}

	message, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil{
		return
	}

	//启用携程来处理函数
	forver := make(chan bool)
	go func(){
		for d:= range message{
			//实现我们要处理的逻辑函数
			log.Printf("resevie message %s",d.Body)
		}
	}()
	log.Printf("[*] Wating for message,to exit press ctrl +c")
	<- forver









}




