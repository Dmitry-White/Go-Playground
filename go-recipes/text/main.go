package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	fmt.Println("------------- Format -------------")
	trades := []Trade{
		{"MSFT", 231, 234.57},
		{"TSLA", 123, 686.75},
		{"BRK-A", 1, 399100},
	}
	genReport(os.Stdout, trades)
	fmt.Println("----------------------------------")

	fmt.Println("------------ Unicode -------------")
	fmt.Println(sentenceLength([]string{"«", "Don't", "Panic", "»"}))
	fmt.Println("----------------------------------")

	fmt.Println("------------- Cases --------------")
	fmt.Println(englishFor("Σ"))
	fmt.Println(englishFor("σ"))
	fmt.Println(englishFor("ς"))
	fmt.Println("----------------------------------")

	fmt.Println("------------- Regex --------------")
	fmt.Printf("%+v\n", parseLedger([]string{"12 shares of MSFT for $234.57"}))
	fmt.Println("----------------------------------")

	fmt.Println("------------- Files --------------")
	fmt.Printf("%#v\n", grep("syslog", "System is rebooting"))

	/* Example history file
	: 1542784278:0;git push
	: 1542784308:0;ls
	: 1542784310:0;go test
	: 1542784314:0;go test -v
	: 1542784386:0;which gometalinter
	: 1542784314:0;go test -v
	*/
	goRegex := regexp.MustCompile(`;go ([a-z]+)`)
	fmt.Println(analyze("zsh_history", goRegex))
	fmt.Println("----------------------------------")
}
