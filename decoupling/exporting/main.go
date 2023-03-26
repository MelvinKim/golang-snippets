package main

type User struct {
	Name     string
	Age      string
	password string
}

type Manager struct {
	Title string
	User
}
