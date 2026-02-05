package main

import "fmt"

func main() {
	fmt.Println("Choose program to run:")
	fmt.Println("1. LRU Cache")
	fmt.Println("2. Task Processor")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		runLRUCache()
	case 2:
		runTaskProcessor()
	default:
		fmt.Println("Invalid choice")
	}
}
