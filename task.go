package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

// Task Processor
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

// Helper Input Functions
func readLine(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func readInt(reader *bufio.Reader, prompt string) int {
	for {
		input := readLine(reader, prompt)
		value, err := strconv.Atoi(input)
		if err == nil {
			return value
		}
		fmt.Println("Please enter a valid number.")
	}
}
func runTaskProcessor() {
	var tasks []Task
	reader := bufio.NewReader(os.Stdin)

	n := readInt(reader, "Enter number of tasks: ")

	for i := 0; i < n; i++ {
		fmt.Println("\nChoose task type:")
		fmt.Println("1. Email Task")
		fmt.Println("2. Data Processing Task")

		choice := readInt(reader, "Enter choice: ")

		switch choice {
		case 1:
			recipient := readLine(reader, "Enter recipient email: ")
			message := readLine(reader, "Enter message: ")

			tasks = append(tasks, EmailTask{
				Recipient: recipient,
				Message:   message,
			})

		case 2:
			data := readLine(reader, "Enter data to process: ")

			tasks = append(tasks, DataProcessingTask{
				Data: data,
			})

		default:
			fmt.Println("Invalid choice, try again")
			i--
		}
	}

	processor := TaskProcessor{Tasks: tasks}
	processor.ProcessTasks()
}
