package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

var ErrNoWorkersProvided = errors.New("no workers provided")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if n <= 0 {
		return ErrNoWorkersProvided
	}
	if len(tasks) == 0 {
		return nil
	}

	tasksChannel := make(chan Task)

	var errorsCount uint32 // Errors count. Simple type used for thread safe atomic.

	wg := sync.WaitGroup{} // Start wait group

	go func() {
		defer close(tasksChannel) // Close channel. There will be no more tasks.
		for _, t := range tasks { // Add task to channel
			if int(atomic.LoadUint32(&errorsCount)) >= m { // Check errors count
				return
			}
			tasksChannel <- t
		}
	}()

	for i := 0; i < n; i++ {
		wg.Add(1) // Add wait for current worker
		go func() {
			defer wg.Done()
			for {
				t, ok := <-tasksChannel
				if !ok {
					break
				}
				if int(atomic.LoadUint32(&errorsCount)) >= m { // Check errors count
					return
				}
				if err := t(); err != nil {
					atomic.AddUint32(&errorsCount, 1) // Increase error count
				}
			}
		}()
	}

	wg.Wait()

	if int(errorsCount) >= m {
		return ErrErrorsLimitExceeded
	}

	return nil
}
