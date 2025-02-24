package flags

import "flag"

var (
	AmqpUrl      string
	AmqpExchange string
	OutputDir    string
)

func init() {
	flag.StringVar(&AmqpUrl, "amqp-url", "amqp://guest:guest@localhost:5672", "URL do RabbitMQ")
	flag.StringVar(&AmqpExchange, "amqp-exchange", "user-events", "Exchange do RabbitMQ")
	flag.StringVar(&OutputDir, "count-out", "./count-out", "Caminho de saída do resumo")
	flag.Parse()
}
