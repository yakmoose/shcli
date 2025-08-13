package server

import (
	"github.com/spf13/cobra"
)

// StackServerCmd represents the cloud command
var Cmd = &cobra.Command{
	Use:   "server",
	Short: "Commands for manipulating with Sitehost cloud stack servers",
}

func init() {
	Cmd.AddCommand(listCmd)
}
