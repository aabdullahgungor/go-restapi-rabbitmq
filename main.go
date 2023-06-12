package main

import (
	"github.com/aabdullahgungor/go-restapi-rabbitmq/server"
)

func main() {

	s := server.NewServer()
	s.Run()
	// var rmqConsumer utils.RMQConsumer
	// rmqConsumer.ConsumeMessage()

}
