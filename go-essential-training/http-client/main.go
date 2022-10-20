package main

import (
	"log"
)

func main() {
	getUsers()

	user := getUser("1")
	log.Printf("User: %+v", user)

	newUser := User{Name: "Test", Username: "TestUsername", Email: "test@test.com"}
	data := addUser(newUser)
	log.Printf("User Data: %+v", data)

	newPost := Post{Title: "Test", UserId: "1", Body: "Test Body"}
	result := addPost(newPost)
	log.Printf("Post Data: %+v", result)

	getPosts()
}
