package ssh

import (
	"shcli/cmd/cloud/ssh/user"

	"github.com/spf13/cobra"
)

// Cmd represents the stacks command
var Cmd = &cobra.Command{
	Use:   "ssh",
	Short: "Manage cloud ssh/sftp",
}

func init() {
	Cmd.AddCommand(user.Cmd)
}
