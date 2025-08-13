package ssh

import (
	"github.com/spf13/cobra"
)

// cloudSSHCmd represents the stacks command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage cloud ssh/sftp users",
}
