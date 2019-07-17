package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// open original file
	originalFile, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()

	// create new file
	newFile, err := os.Create("test-copy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	//copy the bytes from destination to source
	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes", bytesWritten)

	//commit file content
	//flush memory to disk
	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
