package handler

import (
	"github.com/bytebytebug/newf/service"
	"github.com/spf13/cobra"
)

func MakeFileFlags(cmd *cobra.Command) {
	cmd.Flags().String("at", ".", "replace the base path")
}

func MakeFile(cmd *cobra.Command, args []string) error {
	base, err := cmd.Flags().GetString("at")

	if err != nil {
		return err
	}

	makeFileService, err := service.CreateMakeFileService()
	if err != nil {
		return err
	}

	return makeFileService.Exec(base, args)
}
