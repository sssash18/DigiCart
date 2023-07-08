package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sssash18/Digicart/pkg/common/database"
	"github.com/sssash18/Digicart/pkg/common/models"
)

var channelRabbitMQProd *amqp.Channel
var ConnProd *amqp.Connection

func init() {
	database.EnvLoad()
	amqpServerUrl := os.Getenv("AMQP_SERVER_URL")
	fmt.Println(amqpServerUrl)
	conn, err := amqp.Dial(amqpServerUrl)
	if err != nil {
		panic(err)
	}

	channelRabbitMQProd, err = conn.Channel()
	if err != nil {
		panic(err)
	}

	_, err = channelRabbitMQProd.QueueDeclare(
		"NotifyQueue",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

}

func Publish(message *models.Message) error {
	body, _ := json.Marshal(message)
	msg := amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	}
	err := channelRabbitMQProd.PublishWithContext(
		context.Background(),
		"",
		"NotifyQueue",
		false,
		false,
		msg,
	)
	return err
}
