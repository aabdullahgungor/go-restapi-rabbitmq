package utils

import (
	"fmt"
	"net/http"

	"github.com/aabdullahgungor/go-restapi-rabbitmq/models"

	"github.com/aabdullahgungor/go-restapi-rabbitmq/constants"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

type RMQConsumer struct {
	Queue            string
	ConnectionString string
}

func (rmq RMQConsumer) ConsumeMessage(ctx *gin.Context) {
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

	msgs, err := ch.Consume(
		rmq.Queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			if d.Type == constants.TASK1 {
				var productModel models.ProductModel
				fmt.Printf("IN PROGRESS - " + constants.TASK1)
				products, err := productModel.GetAll()
				if err != nil {
					ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "Products cannot show: " + err.Error()})
				}
				ctx.Header("Content-Type", "application/json")
				ctx.IndentedJSON(http.StatusOK, products)
			}
		}
	}()

	fmt.Println("Successfully Connected to our RabbitMQ Instance")
	fmt.Println(" [*] - Waiting for messages")
	<-forever
}
