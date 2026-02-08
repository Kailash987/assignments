# Worker Pool

This module implements a **Worker Pool** pattern in Go to execute multiple tasks concurrently using goroutines and channels, while preserving the order of results.

---

## Overview

The worker pool allows you to:
- Submit multiple independent tasks
- Execute them concurrently using a fixed number of workers
- Collect results in the same order tasks were submitted

---

## Key Components

### Task
Represents a unit of work to be executed.
- Stores a task index
- Holds a function that returns a result

### Result
Stores the output of a completed task along with its index.

### WorkerPool
Manages:
- Number of worker goroutines
- Task submission channel
- Result collection channel
- Task count for ordered output

---

## Workflow

1. Create a worker pool with a fixed number of workers  
2. Submit tasks as functions  
3. Start the pool using `Run()`  
4. Workers process tasks concurrently  
5. Results are collected and returned in submission order  

---

## Concurrency Model

- **Goroutines** are used to run workers concurrently
- **Channels** are used for task distribution and result collection
- **WaitGroup** ensures all workers finish before closing result channels

---

## Ordering Guarantee

Each task is assigned an index at submission time.  
Results are placed into a slice using this index, ensuring output order remains consistent regardless of execution order.

---

## Results:
```
Task 1 → Task 1 processed by worker
Task 2 → Task 2 processed by worker
Task 3 → Task 3 processed by worker
Task 4 → Task 4 processed by worker
Task 5 → Task 5 processed by worker
```
