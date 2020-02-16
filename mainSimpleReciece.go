package main

import (
	"secondskillforgo/rabbitmq"
)


func RecevieMessage(){
	rabbitmqOne := rabbitmq.NewRabbitMQSimple("immocSimple")
	rabbitmqOne.ConsumeSimple()
}