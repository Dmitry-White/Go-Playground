package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Home  string `json:"home"`
	Shell string `json:"shell"`
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

func handleCSVFormat() {
	fmt.Println("[handleCSVFormat] Not Implemented")
	err := errors.New("sample error")
	handleError(err)
}

func handleJSONFormat() {
	fmt.Println("[handleJSONFormat] Not Implemented")
	handleError(nil)
}

func handleOutputFormat() {
	fmt.Println("[handleOutputFormat] Not Implemented")
	handleJSONFormat()
	handleCSVFormat()
}

func handleOutputPath() {
	fmt.Println("[handleOutputPath] Not Implemented")
	handleError(nil)
}

func parseFlags() {
	fmt.Println("[parseFlags] Not Implemented")
	handleError(nil)
}

func collectUsers() []User {
	file, err := os.Open("/etc/passwd")
	handleError(err)
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ':'

	lines, err := reader.ReadAll()
	handleError(err)

	var users []User

	for _, line := range lines {
		id, err := strconv.ParseInt(line[2], 10, 64)
		handleError(err)

		if id < 1000 {
			continue
		}

		user := User{
			Name:  line[0],
			Id:    int(id),
			Home:  line[5],
			Shell: line[6],
		}

		users = append(users, user)
	}

	return users
}

func main() {
	parseFlags()
	users := collectUsers()
	fmt.Println(users)

	handleOutputPath()
	handleOutputFormat()
}
