package main

import (
	"fmt"
	"sync"
	"time"
)

// Task Interface
type Task interface {
	Execute()
}

// Email Task
type EmailTask struct {
	Recipient string
	Message   string
}

func (e EmailTask) Execute() {
	fmt.Printf("Sending email to %s: %s\n", e.Recipient, e.Message)
	time.Sleep(2 * time.Second)
	fmt.Printf("Email sent to %s\n", e.Recipient)
}

// Data Processing Task
type DataProcessingTask struct {
	Data string
}

func (d DataProcessingTask) Execute() {
	fmt.Printf("Processing data: %s\n", d.Data)
	time.Sleep(3 * time.Second)
	fmt.Printf("Data processed: %s\n", d.Data)
}

type TaskProcessor struct {
	Tasks []Task
}

func (tp TaskProcessor) ProcessTasks() {
	var wg sync.WaitGroup

	for _, task := range tp.Tasks {
		wg.Add(1)

		go func(t Task) {
			defer wg.Done()
			t.Execute()
		}(task)
	}

	wg.Wait()
	fmt.Println("All tasks completed.")
}

func runTaskProcessor() {
	tasks := []Task{
		EmailTask{Recipient: "jay@example.com", Message: "Hello Jay"},
		DataProcessingTask{Data: "User analytics"},
		EmailTask{Recipient: "mahi@example.com", Message: "Hello Mahi"},
		DataProcessingTask{Data: "Transaction records"},
	}

	processor := TaskProcessor{Tasks: tasks}
	processor.ProcessTasks()
}
