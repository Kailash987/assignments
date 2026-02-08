# Persistent LRU Cache in Go

## Overview

This project implements a **Least Recently Used (LRU) Cache** in **Go**, enhanced with **file-based persistence**.  
The cache stores key–value pairs, automatically evicts the least recently used item when capacity is exceeded, and saves its state to disk so that it can be restored on the next program run.

---

## Key Features

- **O(1) time complexity** for `GET` and `PUT` operations
- Uses **HashMap + Doubly Linked List** for efficient access and ordering
- **Persistence support** using JSON file storage
- Command-line interaction (`PUT`, `GET`, `EXIT`)
- Automatically restores cache state on restart

---

## Data Structures Used

### 1. Doubly Linked List
- Maintains access order (Most Recently Used → Least Recently Used)
- Allows constant-time insertion and removal

### 2. Hash Map
- Maps keys to linked list nodes
- Enables constant-time lookup

---

## Project Structure

### Core Cache Structures

- `Node`  
  Represents a node in the doubly linked list containing:
  - key
  - value
  - previous and next pointers

- `LRUCache`  
  Main cache structure holding:
  - capacity
  - hash map for fast access
  - head and tail dummy nodes

---


## Core Operations

### `Get(key string) (string, bool)`
- Returns the value if the key exists
- Moves the accessed node to the **Most Recently Used (MRU)** position

### `Put(key string, value string)`
- Inserts or updates a key
- Moves the node to the MRU position
- Evicts the **Least Recently Used (LRU)** item if capacity is exceeded

---


## Persistence Logic

### Saving Cache State

**Method:** `SaveToFile(filename string)`

- Serializes cache data into JSON format
- Stores:
  - Cache capacity
  - Items in MRU → LRU order
- Automatically called before exiting the program

### Loading Cache State

**Method:** `LoadFromFile(filename string)`

- Reads JSON file from disk
- Reconstructs the cache in the correct access order
- Falls back to fresh cache initialization if file is missing

---

## Command-Line Interface

The program supports the following commands:

### `PUT key value`
- Inserts or updates a key-value pair

### `GET key`
- Retrieves the value for a key
- Updates its usage order

### `EXIT`
- Saves cache state to disk
- Safely exits the program

---

## Example Usage

```text
PUT a 10
PUT b 20
GET a
PUT c 30
EXIT
LRU (Least Recently Used) Cache
