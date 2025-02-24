package eventcounter

import (
	"context"
	"math/rand"
	"time"
)

type Consumer interface {
	Created(ctx context.Context, uid string) error
	Updated(ctx context.Context, uid string) error
	Deleted(ctx context.Context, uid string) error
}

// Mutex swappers are used to increase the performance of the event counter making it more concurrent.
type ConsumerWrapper struct {
	CreatedCounts    map[string]int
	createdMuSwapper *MutexSwapper[string]

	UpdatedCounts    map[string]int
	updatedMuSwapper *MutexSwapper[string]

	DeletedCounts    map[string]int
	deletedMuSwapper *MutexSwapper[string]

	consumer Consumer
}

func (c *ConsumerWrapper) randomSleep() {
	time.Sleep(time.Second * time.Duration(rand.Intn(30)))
}

func (c *ConsumerWrapper) Created(ctx context.Context, uid string) error {
	c.createdMuSwapper.Lock(uid)

	searchKeyOrCreateAndIncrement(c.CreatedCounts, uid)

	c.createdMuSwapper.Unlock(uid)

	return nil
}

func (c *ConsumerWrapper) Updated(ctx context.Context, uid string) error {
	c.updatedMuSwapper.Lock(uid)

	searchKeyOrCreateAndIncrement(c.UpdatedCounts, uid)

	c.updatedMuSwapper.Unlock(uid)

	return nil
}

func (c *ConsumerWrapper) Deleted(ctx context.Context, uid string) error {
	c.deletedMuSwapper.Lock(uid)

	searchKeyOrCreateAndIncrement(c.DeletedCounts, uid)

	c.deletedMuSwapper.Unlock(uid)

	return nil
}

func searchKeyOrCreateAndIncrement(m map[string]int, key string) {
	if _, exists := m[key]; !exists {
		m[key] = 1
	} else {
		m[key]++
	}
}

func NewConsumerWrapper() *ConsumerWrapper {
	return &ConsumerWrapper{
		CreatedCounts:    make(map[string]int),
		createdMuSwapper: NewMutexSwapper[string](),

		UpdatedCounts:    make(map[string]int),
		updatedMuSwapper: NewMutexSwapper[string](),

		DeletedCounts:    make(map[string]int),
		deletedMuSwapper: NewMutexSwapper[string](),
	}
}
