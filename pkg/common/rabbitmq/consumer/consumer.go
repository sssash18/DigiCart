package rabbitmq

import (
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sssash18/Digicart/pkg/common/database"
)

var ChannelRabbitMQCons *amqp.Channel
var ConnCons *amqp.Connection

func init() {

	database.EnvLoad()

	amqpServerUrl := os.Getenv("AMQP_SERVER_URL")
	fmt.Println(amqpServerUrl)
	ConnCons, err := amqp.Dial(amqpServerUrl)
	if err != nil {
		panic(err)
	}

	ChannelRabbitMQCons, err = ConnCons.Channel()
	if err != nil {
		panic(err)
	}

	_, err = ChannelRabbitMQCons.QueueDeclare(
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
