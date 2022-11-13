package utils

import (
	"cobraed-cli/core"
	"os"

	"github.com/spf13/cobra"
)

func ShowUsage(cmd *cobra.Command, f *core.Flags) {
	if !f.Prompt && (f.Name == "" || f.Greeting == "") {
		cmd.Usage()
		os.Exit(1)
	}
}
