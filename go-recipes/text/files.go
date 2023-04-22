package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func grep(filename string, term string) interface{} {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	matches := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, term) {
			matches = append(matches, line)
		}
	}

	err = scanner.Err()
	if err != nil {
		return err
	}

	return matches
}

func analyze(filename string, regex *regexp.Regexp) interface{} {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	frequencies := make(map[string]int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		matches := regex.FindStringSubmatch(line)
		if len(matches) == 0 {
			continue
		}

		cmd := matches[1]
		frequencies[cmd]++
	}

	err = scanner.Err()
	if err != nil {
		return err
	}

	return frequencies
}
