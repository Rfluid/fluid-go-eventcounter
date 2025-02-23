Default folder with the outputs of the consumer. After hit `make generator-publish`, you can run the program with

```bash
go run cmd/consumer/main.go -amqp-url="amqp://guest:guest@localhost:5672" -amqp-exchange="user-events" -count-out="./count-out"
```
