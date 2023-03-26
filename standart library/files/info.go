package main

import (
	"fmt"
	"os"
)

// check if a file exists
func checkFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func main() {
	// Use os.Stat to get file stats
	stats, err := os.Stat("./sampleText.txt")
	if err != nil {
		panic(err)
	}

	// get the file's modification time
	fmt.Println("file's modification time: ", stats.ModTime())

	// check the file's mod or permissions
	fmt.Println("file permissions: ", stats.Mode())
	fmode := stats.Mode()
	if fmode.IsRegular() { //check if this is a regular file
		fmt.Println("This is a regular file!")
	}

	// get the file size
	fmt.Println("file size: ", stats.Size())

	// check if it's a directory
	fmt.Println("This is a directory: ", stats.IsDir())

	// check if a file exists
	exists := checkFileExists("./sampleText.txt")
	fmt.Println("file exists check: ", exists)
}
