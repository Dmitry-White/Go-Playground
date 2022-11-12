package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
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

func handleJSONFormat(output *io.Writer, users []User) {
	data, err := json.MarshalIndent(users, "", "  ")
	handleError(err)
	(*output).Write(data)
}

func handleCSVFormat() {
	fmt.Println("[handleCSVFormat] Not Implemented")
	err := errors.New("sample error")
	handleError(err)
}

func handleOutputFormat(output *io.Writer, format string, users []User) {
	if format == "json" {
		handleJSONFormat(output, users)
	} else if format == "csv" {
		handleCSVFormat()
	}
}

func handleOutputPath(path string) io.Writer {
	var output io.Writer

	if path != "" {
		file, err := os.Create(path)
		handleError(err)
		defer file.Close()

		output = file
	} else {
		output = os.Stdin
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

	output := handleOutputPath(path)
	fmt.Println(output)

	handleOutputFormat(&output, format, users)
}
