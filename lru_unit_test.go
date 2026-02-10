package main

import (
	"os"
	"reflect"
	"testing"
)

// CONSTRUCTOR TEST
func TestNewLRUCache(t *testing.T) {
	cache := NewLRUCache(2)

	if cache.capacity != 2 {
		t.Errorf("expected capacity 2, got %d", cache.capacity)
	}

	if cache.head.next != cache.tail {
		t.Error("head should point to tail initially")
	}

	if cache.tail.prev != cache.head {
		t.Error("tail should point back to head initially")
	}

	if len(cache.cache) != 0 {
		t.Error("cache map should be empty")
	}
}

// PUT & GET TESTS

func TestPutAndGet(t *testing.T) {
	cache := NewLRUCache(2)

	cache.Put("a", "1")
	cache.Put("b", "2")

	val, ok := cache.Get("a")
	if !ok || val != "1" {
		t.Errorf("expected value '1', got '%s'", val)
	}
}

func TestGetNonExistentKey(t *testing.T) {
	cache := NewLRUCache(2)

	if _, ok := cache.Get("missing"); ok {
		t.Error("expected key to be missing")
	}
}

// UPDATE EXISTING KEY

func TestPutUpdatesExistingKey(t *testing.T) {
	cache := NewLRUCache(2)

	cache.Put("a", "1")
	cache.Put("a", "updated")

	val, _ := cache.Get("a")
	if val != "updated" {
		t.Errorf("expected updated value, got %s", val)
	}
}

// LRU EVICTION TEST

func TestEviction(t *testing.T) {
	cache := NewLRUCache(2)

	cache.Put("a", "1")
	cache.Put("b", "2")
	cache.Put("c", "3") // should evict "a"

	if _, ok := cache.Get("a"); ok {
		t.Error("expected key 'a' to be evicted")
	}

	if _, ok := cache.Get("b"); !ok {
		t.Error("expected key 'b' to exist")
	}

	if _, ok := cache.Get("c"); !ok {
		t.Error("expected key 'c' to exist")
	}
}

// RECENCY UPDATE TEST

func TestGetMovesToFront(t *testing.T) {
	cache := NewLRUCache(2)

	cache.Put("a", "1")
	cache.Put("b", "2")
	cache.Get("a")      // make "a" MRU
	cache.Put("c", "3") // should evict "b"

	if _, ok := cache.Get("b"); ok {
		t.Error("expected 'b' to be evicted")
	}
}

// INTERNAL remove() TEST

func TestRemove(t *testing.T) {
	cache := NewLRUCache(2)
	cache.Put("a", "1")

	node := cache.cache["a"]
	cache.remove(node)

	if cache.head.next != cache.tail {
		t.Error("node not properly removed from list")
	}
}

// INTERNAL insertAtFront() TEST

func TestInsertAtFront(t *testing.T) {
	cache := NewLRUCache(2)

	node := &Node{key: "x", value: "10"}
	cache.insertAtFront(node)

	if cache.head.next != node {
		t.Error("node was not inserted at front")
	}
}

// evictLRU() TEST

func TestEvictLRU(t *testing.T) {
	cache := NewLRUCache(1)

	cache.Put("a", "1")
	cache.Put("b", "2")

	if _, ok := cache.Get("a"); ok {
		t.Error("expected 'a' to be evicted")
	}
}

// SAVE TO FILE TEST

func TestSaveToFile(t *testing.T) {
	cache := NewLRUCache(2)
	cache.Put("a", "1")
	cache.Put("b", "2")

	filename := "test_cache.json"
	defer os.Remove(filename)

	err := cache.SaveToFile(filename)
	if err != nil {
		t.Fatalf("SaveToFile failed: %v", err)
	}

	if _, err := os.Stat(filename); err != nil {
		t.Error("cache file was not created")
	}
}

// LOAD FROM FILE TEST

func TestLoadFromFile(t *testing.T) {
	cache := NewLRUCache(2)
	cache.Put("a", "1")
	cache.Put("b", "2")

	filename := "test_cache.json"
	defer os.Remove(filename)

	cache.SaveToFile(filename)

	loaded, err := LoadFromFile(filename)
	if err != nil {
		t.Fatalf("LoadFromFile failed: %v", err)
	}

	val, _ := loaded.Get("a")
	if val != "1" {
		t.Errorf("expected '1', got %s", val)
	}
}

// LOAD ORDER PRESERVATION TEST

func TestLoadPreservesOrder(t *testing.T) {
	cache := NewLRUCache(2)
	cache.Put("a", "1")
	cache.Put("b", "2")

	filename := "test_cache.json"
	defer os.Remove(filename)

	cache.SaveToFile(filename)

	loaded, _ := LoadFromFile(filename)

	keys := []string{}
	curr := loaded.head.next
	for curr != loaded.tail {
		keys = append(keys, curr.key)
		curr = curr.next
	}

	expected := []string{"b", "a"} // MRU → LRU
	if !reflect.DeepEqual(keys, expected) {
		t.Errorf("expected order %v, got %v", expected, keys)
	}
}
