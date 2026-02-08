package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// LRU CACHE STRUCTURES

// Node represents a doubly linked list node
type Node struct {
	key   string
	value string
	prev  *Node
	next  *Node
}

// LRUCache represents the cache
type LRUCache struct {
	capacity int
	cache    map[string]*Node
	head     *Node
	tail     *Node
}

// PERSISTENCE STRUCTS

// Struct used to store key-value data in file
type NodeData struct {
	Key   string
	Value string
}

// Struct that represents the entire cache state
type PersistedCache struct {
	Capacity int
	Items    []NodeData
}

// CONSTRUCTOR

func NewLRUCache(capacity int) *LRUCache {
	head := &Node{}
	tail := &Node{}

	head.next = tail
	tail.prev = head

	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*Node),
		head:     head,
		tail:     tail,
	}
}

// CORE OPERATIONS
func (lru *LRUCache) Get(key string) (string, bool) {
	node, ok := lru.cache[key]
	if !ok {
		return "", false
	}

	lru.remove(node)
	lru.insertAtFront(node)

	return node.value, true
}

func (lru *LRUCache) Put(key string, value string) {
	if node, ok := lru.cache[key]; ok {
		node.value = value
		lru.remove(node)
		lru.insertAtFront(node)
		return
	}

	newNode := &Node{
		key:   key,
		value: value,
	}

	lru.cache[key] = newNode
	lru.insertAtFront(newNode)

	if len(lru.cache) > lru.capacity {
		lru.evictLRU()
	}
}
func (lru *LRUCache) remove(node *Node) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
}

func (lru *LRUCache) insertAtFront(node *Node) {
	node.next = lru.head.next
	node.prev = lru.head
	lru.head.next.prev = node
	lru.head.next = node
}

func (lru *LRUCache) evictLRU() {
	lruNode := lru.tail.prev
	lru.remove(lruNode)
	delete(lru.cache, lruNode.key)
}

// DISPLAY
func (lru *LRUCache) Display() {
	fmt.Print("Cache (MRU → LRU): ")
	curr := lru.head.next
	for curr != lru.tail {
		fmt.Printf("[%s:%s] ", curr.key, curr.value)
		curr = curr.next
	}
	fmt.Println()
}

// Saves cache state to a file before exiting
func (lru *LRUCache) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var data PersistedCache
	data.Capacity = lru.capacity

	// traverse cache from MRU → LRU and store items
	curr := lru.head.next
	for curr != lru.tail {
		data.Items = append(data.Items, NodeData{
			Key:   curr.key,
			Value: curr.value,
		})
		curr = curr.next
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

// Loads cache state from file on program start
func LoadFromFile(filename string) (*LRUCache, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data PersistedCache
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}

	cache := NewLRUCache(data.Capacity)

	for i := len(data.Items) - 1; i >= 0; i-- {
		item := data.Items[i]
		cache.Put(item.Key, item.Value)
	}

	return cache, nil
}
func runLRUCache() {
	const filename = "cache.json"

	var cache *LRUCache
	var err error

	cache, err = LoadFromFile(filename)
	if err != nil {
		var capacity int
		fmt.Print("No saved cache found. Enter cache capacity: ")
		fmt.Scan(&capacity)
		cache = NewLRUCache(capacity)
		fmt.Println("Started fresh cache")
	} else {
		fmt.Println("Loaded cache from disk")
		cache.Display()
	}

	fmt.Println("\nCommands:")
	fmt.Println("PUT key value")
	fmt.Println("GET key")
	fmt.Println("EXIT")

	for {
		var command string
		fmt.Print("\nEnter command: ")
		fmt.Scan(&command)

		switch command {
		case "PUT":
			var key, value string
			fmt.Scan(&key, &value)
			cache.Put(key, value)
			fmt.Println("Inserted")
			cache.Display()

		case "GET":
			var key string
			fmt.Scan(&key)
			if val, ok := cache.Get(key); ok {
				fmt.Println("Value:", val)
			} else {
				fmt.Println("Key not found")
			}
			cache.Display()

		case "EXIT":
			cache.SaveToFile(filename)
			fmt.Println("Cache saved. Exiting...")
			return

		default:
			fmt.Println("Invalid command")
		}
	}
}
