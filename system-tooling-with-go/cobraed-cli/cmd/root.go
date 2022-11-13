/*
Copyright © 2022 COBRAED-CLI <dmitry.b.white@gmail.com>
*/
package cmd

import (
	"cobraed-cli/core"
	"cobraed-cli/utils"
	"os"

	"github.com/spf13/cobra"
)

var flags = core.Flags{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cobraed-cli",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.ShowUsage(cmd, &flags)

		utils.HandlePrompt(&flags)

		utils.HandleDebug(&flags)

		utils.HandleMessage(&flags)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	rootCmd.Flags().StringVarP(&flags.Name, "name", "n", "123", "name to use within the message")
	rootCmd.Flags().StringVarP(&flags.Greeting, "greeting", "g", "", "phrase to use within the message")
	rootCmd.Flags().BoolVarP(&flags.Prompt, "prompt", "p", false, "use prompt to input name and greeting")
	rootCmd.Flags().BoolVarP(&flags.Preview, "preview", "v", false, "use preview to output message without writing to a file")
}
