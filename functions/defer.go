package main

import (
	"fmt"
	"os"
)

// opening a file passed as an args in the cmd
func openCmdFile() {
	f := os.Stdin

	if len(os.Args) > 1 {
		if f, err := os.Open(os.Args[1]); err != nil {
			// do something
			fmt.Println(f)
		}

		defer f.Close()
	}
}

func main() {
	f, err := os.OpenFile("notes.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close() //guaranteed to run at function exit

	// write something to the file
	f.WriteString("This is some string\n")
}

/*
 - defers are executed in reverse order i.e from last to first i.e
 func doSomething() {
	 defer a()

	 defer b()
 }
 - defer b() is executed first, then defer a()
*/
