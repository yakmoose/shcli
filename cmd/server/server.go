package server

import (
	"github.com/spf13/cobra"
)

// Cmd represents the cloud command
var Cmd = &cobra.Command{
	Use:   "server",
	Short: "Commands for manipulating with Sitehost servers",
}

func init() {
	Cmd.AddCommand(listCmd)
	Cmd.AddCommand(getCmd)
}
