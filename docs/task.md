# Concurrent Task Processor

## Overview

This project demonstrates a **concurrent task processing system** in **Go** using:

- Interfaces  
- Goroutines  
- `sync.WaitGroup`  
- Command-line input handling  

The program allows users to create different types of tasks and executes them **concurrently**, simulating real-world background job processing.

---

## Key Concepts Demonstrated

- **Interfaces for abstraction**
- **Concurrency using goroutines**
- **Synchronization using WaitGroup**
- **Polymorphism in Go**
- **Command-line interaction**

---

## Architecture Overview

The system follows a simple design:

1. A common `Task` interface defines behavior  
2. Multiple task types implement this interface  
3. A `TaskProcessor` runs all tasks concurrently  
4. A `WaitGroup` ensures the main program waits until all tasks finish  

---


## Concurrency Flow

This section explains how tasks are executed concurrently in the program using goroutines and a `WaitGroup`.

### Execution Steps

1. Loop through all tasks stored in the task list.
2. Increment the `WaitGroup` counter for each task.
3. Launch a new goroutine for every task.
4. Call the `Execute()` method inside the goroutine.
5. Decrement the `WaitGroup` counter once the task completes.
6. Wait until all goroutines finish execution.
7. Print confirmation after all tasks are completed.

### Key Notes

- Each task runs independently.
- Execution order is not guaranteed.
- `WaitGroup` ensures the main program waits for all tasks.

## Example Usage

```text
Enter number of tasks: 2

Choose task type:
1. Email Task
2. Data Processing Task
Enter choice: 1
Enter recipient email: test@example.com
Enter message: Hello

Choose task type:
1. Email Task
2. Data Processing Task
Enter choice: 2
Enter data to process: Dataset A

Sending email to test@example.com: Hello
Processing data: Dataset A
Email sent to test@example.com
Data processed: Dataset A
All tasks completed.
```