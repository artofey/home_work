package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks.
func Run(tasks []Task, N int, M int) error {
	wg := sync.WaitGroup{}
	for _, task := range tasks {
		task = task
		wg.Add(1)
		mu := sync.Mutex{}
		go func() {
			mu.Lock()
			defer mu.Unlock()
			task()
			wg.Done()
		}()
	}

	wg.Wait()
	return nil
}
