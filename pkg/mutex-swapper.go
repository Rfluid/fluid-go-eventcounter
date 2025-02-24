package eventcounter

import (
	"sync"
)

// Swaps a global mutex for a string-specific mutex.
type MutexSwapper[T comparable] struct {
	mutexes               map[T]*sync.Mutex // Map from message id to mutex
	mutexQuantityAcquired map[T]*int8       // Amount of times a mutex was acquired
	Mu                    sync.Mutex        // Global mutex that will be swapped
}

// LockMessage locks the mutex associated with the given message ID.
func (s *MutexSwapper[T]) Lock(msgId T) {
	// Finding mutex and incrementing mutex use must happen
	// inside the global mutex to prevent deleting in use mutexes
	s.Mu.Lock()

	// Finding mutex
	mu, exists := s.mutexes[msgId]
	if !exists {
		mu = &sync.Mutex{}
		s.mutexes[msgId] = mu
	}

	// Incrementing mutext use
	qttAc, exists := s.mutexQuantityAcquired[msgId]
	if !exists {
		var count int8 = 1
		s.mutexQuantityAcquired[msgId] = &count
	} else {
		*qttAc++
	}

	s.Mu.Unlock()

	// Lock the message-specific mutex outside the global mutex to prevent blocking other operations.
	mu.Lock()
}

// UnlockMessage unlocks the mutex associated with the given message ID.
func (s *MutexSwapper[T]) Unlock(msgId T) {
	// Finding mutex and incrementing mutex use must happen
	// inside the global mutex to prevent deleting in use mutexes
	s.Mu.Lock()

	mu, exists := s.mutexes[msgId]
	if !exists {
		// Should not happen; handle error as needed.
		s.Mu.Unlock()
		return
	}

	qttAc, exists := s.mutexQuantityAcquired[msgId]
	if !exists {
		// Should not happen; handle error as needed.
		s.Mu.Unlock()
		return
	}

	*qttAc--
	if *qttAc == 0 {
		delete(s.mutexes, msgId)
		delete(s.mutexQuantityAcquired, msgId)
	}

	s.Mu.Unlock()

	// Unlock the message-specific mutex.
	mu.Unlock()
}

// NewMutexSwapper initializes a new Mutex Swapper.
func NewMutexSwapper[T comparable]() *MutexSwapper[T] {
	return &MutexSwapper[T]{
		mutexes:               make(map[T]*sync.Mutex),
		mutexQuantityAcquired: make(map[T]*int8),
	}
}
