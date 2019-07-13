package main

import (
	"log"
	"os"
)

func main() {
	// pass in 0 to completely empty the file
	err := os.Truncate("test.txt", 100)
	if err != nil {
		log.Fatal(err)
	}
}
