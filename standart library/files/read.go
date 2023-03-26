package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// use the ReadFile function to read an entire file
	_, err := ioutil.ReadFile("sampleText.txt")
	handleError(err)

	// ReadFile reads the data as bytes, we have to convert to string
	// fmt.Println(string(content))

	// it's not always possible to read an entire file into memory at once
	// there will occassions where you need the content in form of chunks
	// for this, you need to opne the file for reading
	// the ReadFile function can perform a generic read
	const BufferSize = 20
	f, err := os.Open("sampleText.txt")
	handleError(err)

	defer f.Close()

	// create a buffer (memory) to read the data into
	b1 := make([]byte, BufferSize)
	for {
		n, err := f.Read(b1)
		if err != nil {
			if err != io.EOF {
				handleError(err)
			}

			break
		}

		fmt.Println("Bytes read: ", n)
		fmt.Println("Content: ", string(b1[:n]))
	}

	// Get the length of a file

	// Use ReadAt to read from a specifc index
}
