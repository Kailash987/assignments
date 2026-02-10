package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

// Test NewWorkerPool
func TestNewWorkerPool(t *testing.T) {
	wp := NewWorkerPool(3)

	if wp.workerCount != 3 {
		t.Errorf("expected workerCount 3, got %d", wp.workerCount)
	}

	if wp.tasks == nil {
		t.Error("tasks channel should not be nil")
	}

	if wp.results == nil {
		t.Error("results channel should not be nil")
	}

	if wp.taskCount != 0 {
		t.Errorf("expected taskCount 0, got %d", wp.taskCount)
	}
}

// Test Submit
func TestSubmit(t *testing.T) {
	wp := NewWorkerPool(2)

	wp.Submit(func() interface{} { return "A" })
	wp.Submit(func() interface{} { return "B" })

	if wp.taskCount != 2 {
		t.Errorf("expected taskCount 2, got %d", wp.taskCount)
	}
}

// Test Run – basic execution
func TestRunBasic(t *testing.T) {
	wp := NewWorkerPool(2)

	wp.Submit(func() interface{} { return "Task 1" })
	wp.Submit(func() interface{} { return "Task 2" })

	results := wp.Run()

	expected := []interface{}{"Task 1", "Task 2"}

	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}

// Test result order preservation
func TestRunPreservesOrder(t *testing.T) {
	wp := NewWorkerPool(3)

	for i := 0; i < 5; i++ {
		taskNum := i
		wp.Submit(func() interface{} {
			return fmt.Sprintf("Task %d", taskNum)
		})
	}

	results := wp.Run()

	for i, r := range results {
		expected := fmt.Sprintf("Task %d", i)
		if r != expected {
			t.Errorf("expected %s at index %d, got %v", expected, i, r)
		}
	}
}

// Test concurrency with delays

func TestRunWithDelays(t *testing.T) {
	wp := NewWorkerPool(3)

	for i := 0; i < 3; i++ {
		taskNum := i
		wp.Submit(func() interface{} {
			time.Sleep(50 * time.Millisecond)
			return taskNum * taskNum
		})
	}

	results := wp.Run()
	expected := []interface{}{0, 1, 4}

	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}

// Test Run with no tasks (edge case)

func TestRunWithNoTasks(t *testing.T) {
	wp := NewWorkerPool(2)

	results := wp.Run()

	if len(results) != 0 {
		t.Errorf("expected empty result slice, got %v", results)
	}
}

// Test single worker

func TestSingleWorker(t *testing.T) {
	wp := NewWorkerPool(1)

	for i := 1; i <= 3; i++ {
		taskNum := i
		wp.Submit(func() interface{} {
			return taskNum
		})
	}

	results := wp.Run()
	expected := []interface{}{1, 2, 3}

	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}
