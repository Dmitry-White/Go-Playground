package main

import (
	"basic-cli/core"
	"basic-cli/utils"
)

func main() {

	flags := core.Flags{}

	utils.ParseFlags(&flags)

	utils.ShowUsage(&flags)

	utils.HandlePrompt(&flags)

	utils.HandleDebug(&flags)

	utils.HandleMessage(&flags)
}
