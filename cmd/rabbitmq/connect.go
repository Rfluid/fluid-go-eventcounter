package rabbitmq

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
	"github.com/reb-felipe/eventcounter/cmd/config/flags"
)

func Connect() {
	// Conecta ao RabbitMQ e consome mensagens da fila "eventcountertest".
	var err error
	Connection, err = amqp091.Dial(flags.AmqpUrl)
	if err != nil {
		log.Fatalf("Erro ao conectar ao RabbitMQ: %v", err)
	}
}
