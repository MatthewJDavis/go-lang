package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := os.Link("original.txt", "original-also.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("creating sym link")
	err = os.Symlink("original.txt", "original-sym.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileInfo, err := os.Lstat("original-sym.txt")
	if fileInfo != nil {
		log.Fatal(err)
	}
	fmt.Printf("Link info '%v'", fileInfo)

	err = os.Lchown("original-sym.txt", os.Getegid(), os.Getegid())
	if err != nil {
		log.Fatal(err)
	}
}
