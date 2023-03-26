package main

import (
	"fmt"
	"os"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func checkFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

// append data to an existing file
// Go does not have an inbuilt append function
// Use the lower-level OpenFile function to open file in append mode
func appendFileData(fname string, data string) {
	f, err := os.OpenFile(fname, os.O_APPEND|os.O_WRONLY, 0644)
	handleError(err)
	defer f.Close()

	// once the file is opened in append mode, we can then add data into it using the regular write function
	_, err = f.WriteString(data)
	handleError(err)
}

func writeFileData() {
	// create a new file
	f, err := os.Create("testfile.txt")
	handleError(err)

	defer f.Close()

	// get the name of the file
	fmt.Println("file name: ", f.Name())

	// write some string to the file
	f.WriteString("This is some string\n")

	// write some individual bytes to the file
	data2 := []byte("writting to a file using a byte array")
	f.Write(data2)

	// the Sync function forces data to be written out to Disk
	f.Sync()
}

func main() {

	// Granular writing of data to a file --> when the write operation doesn't happen at once
	// happens at different stages
	// truncating a file == changing size of a file
	if !checkFileExists("testfile.txt") {
		writeFileData()
	} else {
		os.Truncate("testfile.txt", 20)
		fmt.Println("Trimmed the file data")
	}

	// append data to a file
	appendFileData("datafile.txt", "This data was appended\n")

	// dump some data to a file --> must be inform of []bytes
	/*
		data1 := []byte("This is some text data\n")
		if err := ioutil.WriteFile("datafile.txt", data1, 0644); err != nil {
			fmt.Println("error while writing data to a file")
		}
		fmt.Println("data written successfully")
	*/
}
