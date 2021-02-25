package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n int, m int) error {
	wg := sync.WaitGroup{}
	doneCh := make(chan interface{})
	tasksCh := make(chan Task, len(tasks))
	errorsCh := make(chan error)
	moreError := false

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			worker(doneCh, tasksCh, errorsCh)
			wg.Done()
		}()
	}

	for _, task := range tasks {
		tasksCh <- task
	}
	close(tasksCh)

	var errors int
	for i := 0; i < len(tasks); i++ {
		err := <-errorsCh
		if err != nil {
			errors++
			if errors >= m {
				close(doneCh)
				moreError = true
				break
			}
		}
	}

	wg.Wait()
	close(errorsCh)
	if moreError {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func worker(done <-chan interface{}, tasks <-chan Task, errors chan<- error) {
	for task := range tasks {
		err := task()
		select {
		case <-done:
			return
		case errors <- err:
		}
	}
}
