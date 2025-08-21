package key

import (
	"github.com/spf13/cobra"
)

// Cmd represents the stacks command
var Cmd = &cobra.Command{
	Use:   "key",
	Short: "Manage ssh keys",
}
