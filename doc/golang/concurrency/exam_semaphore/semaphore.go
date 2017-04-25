package main

import (
	"errors"
	"sync"
	"time"
)

// Semaphore is an implementation of semaphore.
type Semaphore struct {
	permits int
	avail   int
	channel chan struct{}
	rwMutex *sync.RWMutex
}

// New creates a new Semaphore with specified number of permits.
func New(permits int) *Semaphore {
	if permits < 1 {
		panic("Invalid number of permits. Less than 1")
	}

	// fill channel buffer
	channel := make(chan struct{}, permits)
	for i := 0; i < permits; i++ {
		channel <- struct{}{}
	}

	return &Semaphore{
		permits: permits,
		avail:   permits,
		channel: channel,
		rwMutex: &sync.RWMutex{},
	}
}

// Acquire acquires one permit. If it is not available, the goroutine will block until it is available.
func (s *Semaphore) Acquire() (bool, error) {
	<-s.channel

	s.rwMutex.Lock()
	s.avail--
	s.rwMutex.Unlock()

	return true, nil
}

// AcquireWithin is similar to AcquireMany() but cancels if duration elapses before getting the permits.
// Returns true if successful and false if timeout occurs.
// timeout unit ms
func (s *Semaphore) AcquireTimeout(timeout int) (bool, error) {
	timeoutChan := make(chan bool, 1)
	cancelChan := make(chan bool, 1)

	go func() {
		time.Sleep(time.Millisecond * time.Duration(timeout))
		timeoutChan <- true
	}()

	go func() {
		s.Acquire()

		timeoutChan <- false
		if <-cancelChan {
			s.Release()
		}
	}()

	if <-timeoutChan {
		cancelChan <- true
		return false, errors.New("timeout occurs")
	}
	cancelChan <- false

	return true, nil
}

// Release releases one permit.
func (s *Semaphore) Release() {
	s.rwMutex.RLock()
	if s.avail == s.permits {
		s.rwMutex.RUnlock()
		return
	}
	s.rwMutex.RUnlock()

	s.channel <- struct{}{}

	s.rwMutex.Lock()
	s.avail++
	s.rwMutex.Unlock()
}

// AvailablePermits gives number of available unacquired permits.
func (s *Semaphore) AvailablePermits() int {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()

	if s.avail < 0 {
		return 0
	}

	return s.avail
}
