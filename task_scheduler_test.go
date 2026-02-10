package main

import (
	"bufio"
	"bytes"
	"os"
	"strings"
	"testing"
)

// Fake task for testing
type FakeTask struct {
	Executed bool
}

func (f *FakeTask) Execute() {
	f.Executed = true
}

// Email task test
func TestEmailTaskExecute(t *testing.T) {
	var output bytes.Buffer

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	task := EmailTask{
		Recipient: "test@example.com",
		Message:   "Hello",
	}

	go func() {
		task.Execute()
		w.Close()
	}()

	output.ReadFrom(r)
	os.Stdout = oldStdout

	result := output.String()

	if !strings.Contains(result, "Sending email to test@example.com") {
		t.Errorf("Missing send output")
	}

	if !strings.Contains(result, "Email sent to test@example.com") {
		t.Errorf("Missing sent output")
	}
}

// Data task test
func TestDataProcessingTaskExecute(t *testing.T) {
	var output bytes.Buffer

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	task := DataProcessingTask{
		Data: "SampleData",
	}

	go func() {
		task.Execute()
		w.Close()
	}()

	output.ReadFrom(r)
	os.Stdout = oldStdout

	result := output.String()

	if !strings.Contains(result, "Processing data: SampleData") {
		t.Errorf("Missing processing output")
	}

	if !strings.Contains(result, "Data processed: SampleData") {
		t.Errorf("Missing processed output")
	}
}

// Processor test
func TestTaskProcessorProcessTasks(t *testing.T) {
	task1 := &FakeTask{}
	task2 := &FakeTask{}

	processor := TaskProcessor{
		Tasks: []Task{task1, task2},
	}

	processor.ProcessTasks()

	if !task1.Executed || !task2.Executed {
		t.Errorf("Tasks not executed")
	}
}

// readLine test
func TestReadLine(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("hello\n"))

	result := readLine(reader, "Input: ")

	if result != "hello" {
		t.Errorf("Unexpected value")
	}
}

// readInt valid input
func TestReadIntValid(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("5\n"))

	result := readInt(reader, "Number: ")

	if result != 5 {
		t.Errorf("Expected 5")
	}
}

// readInt invalid then valid
func TestReadIntInvalidThenValid(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("x\n7\n"))

	result := readInt(reader, "Number: ")

	if result != 7 {
		t.Errorf("Expected 7")
	}
}
