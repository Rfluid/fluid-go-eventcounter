package exchange_listener

import (
	"encoding/json"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/reb-felipe/eventcounter/cmd/rabbitmq"
	eventcounter "github.com/reb-felipe/eventcounter/pkg"
)

// Connects to RabbitMQ and consumes messages from the exchange.
func ConsumeLoop() {
	// Conecta ao RabbitMQ e garante que a conexão será fechada ao final.
	rabbitmq.Connect()
	defer rabbitmq.Connection.Close()

	// Abre um canal de comunicação com o RabbitMQ e garante que o canal será fechado ao final.
	ch, err := rabbitmq.Connection.Channel()
	if err != nil {
		log.Fatalf("Erro ao abrir canal: %v", err)
	}
	defer ch.Close()

	// Consome mensagens (auto-ack para simplificar)
	msgs, err := ch.Consume("eventcountertest", "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Erro ao consumir mensagens: %v", err)
	}

	// Mapa para garantir que cada mensagem (identificada por seu ID) seja processada apenas uma vez.
	processed := make(map[string]bool)
	var procMu sync.Mutex

	// Timer para detectar 5 segundos de inatividade.
	timer := time.NewTimer(5 * time.Second)

consumeLoop:
	for {
		select {
		case d, ok := <-msgs:
			if !ok {
				log.Println("Canal de mensagens fechado")
				break consumeLoop
			}

			// Reinicia o timer de inatividade.
			if !timer.Stop() {
				select {
				case <-timer.C:
				default:
				}
			}
			timer.Reset(5 * time.Second)

			// A routing key segue o formato: <user_id>.event.<tipo do evento>
			parts := strings.Split(d.RoutingKey, ".")
			if len(parts) != 3 {
				log.Printf("Routing key inválida: %s", d.RoutingKey)
				continue
			}

			userID := parts[0]
			eventType := eventcounter.EventType(parts[2])

			// O corpo da mensagem é um JSON com {"id": "id_unico"}
			var body struct {
				ID string `json:"id"`
			}
			if err := json.Unmarshal(d.Body, &body); err != nil {
				log.Printf("Erro ao fazer unmarshal da mensagem: %v", err)
				continue
			}

			// Verifica se a mensagem já foi processada.
			procMu.Lock()
			if processed[body.ID] {
				procMu.Unlock()
				continue
			}
			processed[body.ID] = true
			procMu.Unlock()

			// Encaminha a mensagem para o canal correspondente.
			em := eventcounter.Message{
				UserID:    userID,
				EventType: eventType,
				UID:       body.ID,
			}
			switch eventType {
			case "created":
				createdCh <- em
			case "updated":
				updatedCh <- em
			case "deleted":
				deletedCh <- em
			default:
				log.Printf("Tipo de evento desconhecido: %s", eventType)
			}

		case <-timer.C:
			log.Println("Inatividade de 5 segundos detectada. Encerrando consumo.")
			break consumeLoop
		}
	}

	// Fecha os canais para os workers e aguarda o término dos processamentos.
	close(createdCh)
	close(updatedCh)
	close(deletedCh)
	wg.Wait()
}
