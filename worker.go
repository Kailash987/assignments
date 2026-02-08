package main

import (
	"fmt"
	"sync"
)

//Internal Structures

type task struct {
	index int
	work  func() interface{}
}

type result struct {
	index int
	value interface{}
}

//WorkerPool Definition

type WorkerPool struct {
	workerCount int
	tasks       chan task
	results     chan result
	taskCount   int
}

func NewWorkerPool(workerCount int) *WorkerPool {
	return &WorkerPool{
		workerCount: workerCount,
		tasks:       make(chan task, 100),
		results:     make(chan result),
	}
}

//Submit Task

func (wp *WorkerPool) Submit(t func() interface{}) {
	wp.tasks <- task{
		index: wp.taskCount,
		work:  t,
	}
	wp.taskCount++
}

//Run Worker Pool

func (wp *WorkerPool) Run() []interface{} {
	var wg sync.WaitGroup

	for i := 0; i < wp.workerCount; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for t := range wp.tasks {
				resultValue := t.work()
				wp.results <- result{
					index: t.index,
					value: resultValue,
				}
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(wp.results)
	}()

	close(wp.tasks)

	output := make([]interface{}, wp.taskCount)
	for res := range wp.results {
		output[res.index] = res.value
	}

	return output
}

func runWorkerPool() {
	wp := NewWorkerPool(3)

	for i := 1; i <= 5; i++ {
		taskNum := i
		wp.Submit(func() interface{} {
			return fmt.Sprintf("Task %d completed", taskNum)
		})
	}

	results := wp.Run()

	fmt.Println("Results:")
	for i, r := range results {
		fmt.Printf("Task %d → %v\n", i+1, r)
	}
}
