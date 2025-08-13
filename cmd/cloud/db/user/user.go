package user

import (
	"github.com/spf13/cobra"
)

// Cmd represents the db user commands
var Cmd = &cobra.Command{
	Use:   "user",
	Short: "Manage cloud database users",
}

func init() {
	Cmd.AddCommand(addCmd)
	Cmd.AddCommand(attachCmd)
	Cmd.AddCommand(deleteCmd)
	Cmd.AddCommand(detachCmd)
	Cmd.AddCommand(getCmd)

	Cmd.AddCommand(listCmd)
	Cmd.AddCommand(updateCmd)
}
