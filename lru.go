package main

import "fmt"

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

// Constructor
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

// Get retrieves value and updates usage
func (lru *LRUCache) Get(key string) (string, bool) {
	node, ok := lru.cache[key]
	if !ok {
		return "", false
	}

	lru.remove(node)
	lru.insertAtFront(node)

	return node.value, true
}

// Put inserts or updates value
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

// Remove node from list
func (lru *LRUCache) remove(node *Node) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
}

// Insert node right after head
func (lru *LRUCache) insertAtFront(node *Node) {
	node.next = lru.head.next
	node.prev = lru.head
	lru.head.next.prev = node
	lru.head.next = node
}

// Evict least recently used node
func (lru *LRUCache) evictLRU() {
	lruNode := lru.tail.prev
	lru.remove(lruNode)
	delete(lru.cache, lruNode.key)
}

// Display cache state (optional but useful)
func (lru *LRUCache) Display() {
	fmt.Print("Cache (MRU → LRU): ")
	curr := lru.head.next
	for curr != lru.tail {
		fmt.Printf("[%s:%s] ", curr.key, curr.value)
		curr = curr.next
	}
	fmt.Println()
}

func runLRUCache() {
	var capacity int
	fmt.Print("Enter cache capacity: ")
	fmt.Scan(&capacity)

	cache := NewLRUCache(capacity)

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
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid command")
		}
	}

}
