package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Println("User email: ", u.email)
}

type admin struct {
	user  //Embedded Type
	level string
}

// func (a *admin) notify() {
// 	fmt.Printf("Sending admin Email To %s<%s>\n", a.name, a.email)
// }

func main() {
	ad := admin{
		user: user{
			name:  "John Smith",
			email: "john@gmail.com",
		},
		level: "super",
	}

	// we can access the inner type's method directly
	ad.user.notify()

	// The inner type's method is promoted
	ad.notify()

	// send the admin user a notification
	// The embedded inner type's implementation of the
	// interface is "promoted" to the outer type
	sendNotification(&ad)
}

// Accepts values that implement the notifier interface and send notifications
func sendNotification(n notifier) {
	n.notify()
}
