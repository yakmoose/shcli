package ssh

import (
	"github.com/spf13/cobra"
)

// stackCommand represents the stacks command
var sshKeysCmd = &cobra.Command{
	Use:   "key",
	Short: "Manage ssh keys",
}

func init() {
	Cmd.AddCommand(sshKeysCmd)
}
