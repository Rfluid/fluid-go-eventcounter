package exchange_listener

import (
	"sync"

	eventcounter "github.com/reb-felipe/eventcounter/pkg"
)

// Inicializa o contador de eventos.
var (
	Cw = eventcounter.NewConsumerWrapper() // Usado para contar eventos.
	wg sync.WaitGroup                      // Usado para garantir que todas as mensagens sejam processadas antes de finalizar o programa. Adições acontecem na função Wait e a espera acontece na função ConsumeLoop.
)

// Cria canais para cada tipo de evento.
var (
	createdCh = make(chan eventcounter.Message, 100)
	updatedCh = make(chan eventcounter.Message, 100)
	deletedCh = make(chan eventcounter.Message, 100)
)
