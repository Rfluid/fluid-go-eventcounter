package main

import (
	exchange_listener "github.com/reb-felipe/eventcounter/cmd/exchange-listener"
	"github.com/reb-felipe/eventcounter/cmd/out"
)

func main() {
	exchange_listener.Wait()
	exchange_listener.ConsumeLoop()
	out.GenerateConsumerOutputs()
}
