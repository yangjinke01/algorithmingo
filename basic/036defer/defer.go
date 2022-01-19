package main

import (
	"fmt"
	"os"
)

func createFile(path string) *os.File {
	fmt.Println("creating")
	result, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return result
}

func writeFile(file *os.File, content *string) {
	fmt.Println("writing")
	fmt.Fprintln(file, *content)
}

func closeFile(file *os.File) {
	fmt.Println("closing")
	err := file.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
func main() {
	filepath := "/tmp/file.txt"
	content := "happy everyday"
	file := createFile(filepath)
	defer closeFile(file)
	writeFile(file, &content)
}
