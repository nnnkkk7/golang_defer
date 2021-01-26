package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f := createFile("defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(s string) *os.File {
	fmt.Println("create")
	f, err := os.Create(s)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func closeFile(f *os.File) {
	fmt.Println("close")
	err := f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func writeFile(f *os.File) {
	fmt.Println("write")
	fmt.Fprintln(f, "defer example")
}
