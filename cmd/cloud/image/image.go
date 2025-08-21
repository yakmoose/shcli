package image

import (
	"github.com/spf13/cobra"
)

// Cmd represents the stacks command
var Cmd = &cobra.Command{
	Use:   "image",
	Short: "Manage cloud images",
}

func init() {
	Cmd.AddCommand(listCmd)
}
