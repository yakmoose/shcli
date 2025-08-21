package srs

import (
	"shcli/cmd/srs/contacts"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "domain",
	Short: "Commands managing domain names",
}

func init() {
	Cmd.AddCommand(listCmd)
	Cmd.AddCommand(contacts.Cmd)
}
