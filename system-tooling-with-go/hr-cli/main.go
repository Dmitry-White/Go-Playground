package main

import (
	"encoding/csv"
	"encoding/json"
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

func handleJSONFormat(output *os.File, users []User) {
	data, err := json.MarshalIndent(users, "", "  ")
	handleError(err)
	output.Write(data)
	output.Close()
}

func handleCSVFormat(output *os.File, users []User) {
	output.Write([]byte("name,id,home,shell\n"))

	writer := csv.NewWriter(output)
	for _, user := range users {
		err := writer.Write([]string{user.Name, strconv.Itoa(user.Id), user.Home, user.Shell})
		handleError(err)
	}

	writer.Flush()
	output.Close()
}

func handleOutputFormat(output *os.File, format string, users []User) {
	if format == "json" {
		handleJSONFormat(output, users)
	} else if format == "csv" {
		handleCSVFormat(output, users)
	}
}

func handleOutputPath(path string) *os.File {
	var output *os.File

	if path != "" {
		file, err := os.Create(path)
		handleError(err)

		output = file
	} else {
		output = os.Stdout
	}

	return output
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
	file, err := os.Open("passwd")
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
	fmt.Println(format)

	users := collectUsers()

	output := handleOutputPath(path)

	handleOutputFormat(output, format, users)
}
