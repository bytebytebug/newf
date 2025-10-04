package cmd

import (
	"os"

	"github.com/bytebytebug/newf/handler"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "newf",
	Short: "An alternative to touch",
	Long: `An alternative to the touch command.
The main purpose of this command is to create files in a simpler way, without having to worry about existing folders or files.`,
	RunE: handler.MakeFile,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	handler.MakeFileFlags(rootCmd)
}
