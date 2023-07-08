package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/sssash18/Digicart/pkg/common/database"
	rabbitmqCons "github.com/sssash18/Digicart/pkg/common/rabbitmq/consumer"
)

func main() {

	database.EnvLoad()

	messages, err := rabbitmqCons.ChannelRabbitMQCons.Consume(
		"NotifyQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening for messages on Notify Queue")

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for msg := range messages {
			fmt.Println("Received message")
			Notifier(msg.Body)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()
	<-sigCh
}
