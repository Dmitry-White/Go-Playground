package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func parseFlags() (path, format string) {
	flag.StringVar(&path, "path", "", "the path to export file.")
	flag.StringVar(&format, "format", "json", "the output format for the user information. Available options are 'csv' and 'json'. The default option is 'json'.")
	flag.Parse()

	format = strings.ToLower(format)

	if format != "csv" && format != "json" {
		flag.Usage()
		err := errors.New("ivalid format. Use 'json' or 'csv' instead")
		handleError(err)
	}

	return
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
	path, format := parseFlags()
	fmt.Println(path, format)

	users := collectUsers()
	fmt.Println(users)

	handleOutputPath()
	handleOutputFormat()
}
