package main

type Todo struct {
	Completed bool
	Id        int
	Date      string
	Todo      string
}

type Note struct {
	Title   string
	Id      int
	Date    string
	Content string
}

type User struct {
	Name     string
	Email    string
	Password string
}
