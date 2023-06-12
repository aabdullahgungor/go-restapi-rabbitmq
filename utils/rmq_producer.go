package utils

import (
	"fmt"

	"github.com/aabdullahgungor/go-restapi-rabbitmq/constants"
	"github.com/streadway/amqp"
)

type RMQProducer struct {
	Queue            string
	ConnectionString string
}

func (rmq RMQProducer) PublishMessage(contentType string, body []byte, task string) error {
	rmq.Queue = constants.QUEUE
	rmq.ConnectionString = fmt.Sprintf("amqp://%s:%s@%s:%s/", constants.USERNAME, constants.PASSWORD, constants.HOST, constants.PORT)

	conn, errConn := amqp.Dial(rmq.ConnectionString)
	if errConn != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(errConn)
	}

	ch, errCh := conn.Channel()
	if errCh != nil {
		fmt.Println(errCh)
	}
	defer ch.Close()

	q, errQ := ch.QueueDeclare(
		rmq.Queue, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	fmt.Println(q)

	if errQ != nil {
		fmt.Println(errQ)
	}

	errPub := ch.Publish(
		"",        // exchange
		rmq.Queue, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: contentType,
			Body:        body,
			Type:        task,
		},
	)

	if errPub != nil {
		fmt.Println(errPub)
	}
	fmt.Println("Successfully Published Message to Queue")
	return nil
}
