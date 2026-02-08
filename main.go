package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nChoose program to run:")
		fmt.Println("1. LRU Cache")
		fmt.Println("2. Task Processor")
		fmt.Println("3. Worker Pool")
		fmt.Println("4. Exit")

		choice := readInt(reader, "Enter choice: ")

		switch choice {
		case 1:
			runLRUCache()
		case 2:
			runTaskProcessor()
		case 3:
			runWorkerPool()
		case 4:
			fmt.Println("Exiting main program...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
