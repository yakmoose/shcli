package contacts

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "contact",
	Short: "Commands managing domain name contacts",
}

func init() {
	Cmd.AddCommand(listCmd)
}
