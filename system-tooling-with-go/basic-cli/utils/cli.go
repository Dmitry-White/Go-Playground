package utils

import (
	"basic-cli/core"
	"flag"
	"os"
)

func ParseFlags(f *core.Flags) {
	flag.StringVar(&f.Name, "name", "123", "name to use within the message")
	flag.StringVar(&f.Greeting, "greeting", "", "phrase to use within the message")
	flag.BoolVar(&f.Prompt, "prompt", false, "use prompt to input name and greeting")
	flag.BoolVar(&f.Preview, "preview", false, "use preview to output message without writing to a file")

	flag.Parse()
}

func ShowUsage(f *core.Flags) {
	// Note:
	// f.Prompt is a syntactic sugar for
	// (*f).Name
	// to reduce typing
	if !f.Prompt && (f.Name == "" || f.Greeting == "") {
		flag.Usage()
		os.Exit(1)
	}
}
