package main

type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Post struct {
	Title  string `json:"title"`
	UserId string `json:"userId"`
	Body   string `json:"body"`
}
