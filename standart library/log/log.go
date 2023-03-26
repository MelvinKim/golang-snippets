package main

import "log"

func init() {
	log.SetPrefix("TRACE: ")
	// These Flags control the information that are printed in the logs
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	// writes message to the standard logger
	log.Println("message")

	// writes message to the standard logger and calls os.Exit
	log.Fatalln("fatal message")

	// writes message to the standard logger and calls panic()
	log.Panicln("panic message")
}

/*
- init() is executed before main()
- this happens during program initialization
*/
