package main

import (
	"fmt"
	"os"
)

func main() {
	fileData, err := os.ReadFile("/tmp/file")
	checkError(err)
	fmt.Println(string(fileData))

	bytes := make([]byte, 5)
	file, err := os.Open("/tmp/file")
	checkError(err)
	count, err := file.Read(bytes)
	checkError(err)
	fmt.Printf("%d bytes: %s\n", count, bytes[:count])

	sfile, err := file.Seek(6, 0)
	checkError(err)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
