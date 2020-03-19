// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Package task provides a pool of goroutines to perform tasks.
package task

import "sync"

// Worker must be implemented by types that want to use
// the run pool.
type Worker interface {
	Work() string
}

// Task provides a pool of goroutines that can execute any Worker
// tasks that are submitted.
type Task struct {
	work chan Worker
	wg   sync.WaitGroup
	Res  chan string
}

// New creates a new work pool.
func New(maxGoroutines int) *Task {
	t := Task{

		// Using an unbuffered channel because we want the
		// guarantee of knowing the work being submitted is
		// actually being worked on after the call to Run returns.
		work: make(chan Worker),
		Res:  make(chan string, 3),
	}

	// The goroutines are the pool. So we could add code
	// to change the size of the pool later on.

	t.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range t.work {
				r := w.Work()
				t.Res <- r
			}
			t.wg.Done()
		}()
	}

	return &t
}

// Shutdown waits for all the goroutines to shutdown.
func (t *Task) Shutdown() {
	close(t.work)
	close(t.Res)
	t.wg.Wait()
}

// Do submits work to the pool.
func (t *Task) Do(w Worker) {
	t.work <- w
}
