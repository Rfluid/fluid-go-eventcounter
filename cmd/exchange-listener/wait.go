package exchange_listener

import (
	"context"
	"log"
)

func Wait() {
	// WaitGroup para aguardar os workers dos canais.
	wg.Add(3)
	ctx := context.Background()

	// Worker para eventos "created"
	go func() {
		defer wg.Done()
		for msg := range createdCh {
			if err := Cw.Created(ctx, msg.UserID); err != nil {
				log.Printf("Erro processando created para %s: %v", msg.UserID, err)
			}
		}
	}()

	// Worker para eventos "updated"
	go func() {
		defer wg.Done()
		for msg := range updatedCh {
			if err := Cw.Updated(ctx, msg.UserID); err != nil {
				log.Printf("Erro processando updated para %s: %v", msg.UserID, err)
			}
		}
	}()

	// Worker para eventos "deleted"
	go func() {
		defer wg.Done()
		for msg := range deletedCh {
			if err := Cw.Deleted(ctx, msg.UserID); err != nil {
				log.Printf("Erro processando deleted para %s: %v", msg.UserID, err)
			}
		}
	}()
}
