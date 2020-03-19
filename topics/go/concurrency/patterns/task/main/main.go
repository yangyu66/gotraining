// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This sample program demonstrates how to use the work package
// to use a pool of goroutines to get work done.
package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/ardanlabs/gotraining/topics/go/concurrency/patterns/task"
)

// names provides a set of names need to display / print.
var printnametasks = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
	"chen",
	"bob",
	"mary",
	"therese",
	"jason",
	"chen",
}

// namePrinter provides special support for printing names.
type namePrinter struct {
	name string
}

// Work implements the Worker interface.
func (m namePrinter) Work() string {
	log.Println(m.name)
	time.Sleep(2 * time.Second)
	return m.name
}

func main() {
	const routines = 5

	// Create a task pool.
	t := task.New(routines)

	var wg sync.WaitGroup
	wg.Add(len(printnametasks))

	// Iterate over the slice of names.
	for _, name := range printnametasks {

		// Create a namePrinter and provide the
		// specific name.
		np := namePrinter{
			name: name,
		}

		go func() {

			// Submit the task to be worked on. When Do
			// returns, we know it is being handled.
			t.Do(np)
			result := <-t.Res
			fmt.Println(result)
			wg.Done()
		}()
	}

	wg.Wait()

	// Shutdown the task pool and wait for all existing work
	// to be completed.
	t.Shutdown()
}
