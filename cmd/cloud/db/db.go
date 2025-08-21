package db

import (
	"shcli/cmd/cloud/db/grant"
	"shcli/cmd/cloud/db/user"

	"github.com/spf13/cobra"
)

// Cmd represents the stacks command
var Cmd = &cobra.Command{
	Use:   "db",
	Short: "Manage cloud databases",
}

func init() {
	Cmd.AddCommand(addCmd)
	Cmd.AddCommand(rmCmd)
	Cmd.AddCommand(listCmd)
	Cmd.AddCommand(updateCmd)
	Cmd.AddCommand(getCmd)

	Cmd.AddCommand(user.Cmd)
	Cmd.AddCommand(grant.Cmd)
}
