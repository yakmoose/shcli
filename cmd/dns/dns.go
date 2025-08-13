package dns

import (
	"github.com/spf13/cobra"
)

// md represents the domain command
var Cmd = &cobra.Command{
	Use:   "dns",
	Short: "Commands for managing dns zones",
}

func init() {

	Cmd.AddCommand(listCmd)
	Cmd.AddCommand(deleteCmd)
	Cmd.AddCommand(addCmd)

	Cmd.AddCommand(RecordCmd)
}
