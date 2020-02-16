package main


import (
	"secondskillforgo/rabbitmq"
)

func main(){
	rabbitmqOne := rabbitmq.NewRabbitMQSimple("immocSimple")
	rabbitmqOne.PublishSimple("Hello immoc")

}