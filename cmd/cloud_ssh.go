package cmd

import (
	"github.com/spf13/cobra"
)

// cloudSSHCmd represents the stacks command
var cloudSSHCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Manage cloud ssh/sftp",
}

func init() {
	cloudCommand.AddCommand(cloudSSHCmd)
}
