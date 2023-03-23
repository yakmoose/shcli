package cmd

import (
	"github.com/spf13/cobra"
)

// cloudSSHCmd represents the stacks command
var cloudSSHUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage cloud ssh/sftp users",
}

func init() {
	cloudSSHCmd.AddCommand(cloudSSHUserCmd)
}
