

# Assignments 
This repository contains multiple Go programs demonstrating **core Golang concepts** such as data structures, concurrency, goroutines, interfaces, and worker pools.  
All programs are combined under a single entry point (`main.go`) with a menu-driven interface.

---


##  Programs Included

### 1. LRU Cache (lru.go)
- Implements an **LRU (Least Recently Used) Cache**
- Uses:
  - HashMap for O(1) lookup
  - Doubly Linked List for maintaining usage order
- Supports:
  - `PUT key value`
  - `GET key`
- Automatically evicts the least recently used item when capacity is exceeded

**Concepts used:**
- Structs
- Pointers
- Doubly linked list
- Maps
- Encapsulation using methods

---

### 2 Concurrent Task Processor (task.go)
- Processes different tasks concurrently using **goroutines**
- Task types:
  - Email Task
  - Data Processing Task
- Uses an interface-based design so new task types can be added easily

**Concepts used:**
- Interfaces
- Goroutines
- WaitGroups
- Concurrency patterns

---

### 3 Worker Pool (worker.go)
- Implements a **Worker Pool pattern**
- Fixed number of workers execute submitted tasks concurrently
- Maintains task order using task indexing
- Collects results safely

**Concepts used:**
- Channels
- Goroutines
- Synchronization
- Worker pool design pattern

---

##  How to Run

Step 1: Initialize module (if not already)
```bash
go mod init assignments-main
````

### Step 2: Run the program

```bash
go run .
```

OR

```bash
go run main.go lru.go task.go worker.go
```

---

## 🧭 Program Flow

When you run the program, you will see:

```
Choose program to run:
1. LRU Cache
2. Task Processor
3. Worker Pool
4. Exit
```

## 🛠 Requirements

* Go 1.20+ (recommended)
* Basic understanding of Go syntax

---

