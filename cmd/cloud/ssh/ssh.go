package ssh

import (
	"github.com/spf13/cobra"
)

// Cmd represents the stacks command
var Cmd = &cobra.Command{
	Use:   "ssh",
	Short: "Manage cloud ssh/sftp",
}

func init() {
	Cmd.AddCommand(userCmd)
}
