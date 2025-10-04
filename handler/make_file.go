package handler

import (
	"fmt"

	"github.com/spf13/cobra"
)

func MakeFileFlags(cmd *cobra.Command) {
	cmd.Flags().String("at", ".", "replace the base path")
}

func MakeFile(cmd *cobra.Command, args []string) error {
	base, err := cmd.Flags().GetString("at")

	fmt.Println(base)
	fmt.Println(args)

	return err
}
