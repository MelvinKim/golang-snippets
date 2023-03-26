package main

import "fmt"

type printer interface {
	print()
}

type list []printer

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Sorry. We are waiting for delivery :)")
		return
	}

	for _, it := range l {
		it.print()
	}
}
