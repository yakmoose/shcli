package environment

import (
	onepassword "shcli/cmd/cloud/stack/environment/op"

	"github.com/spf13/cobra"
)

// Cmd represents the domain command
var Cmd = &cobra.Command{
	Use:   "env",
	Short: "Commands for managing stack environment variables",
}

func init() {

	Cmd.AddCommand(onepassword.Cmd)

	Cmd.AddCommand(getCmd)
	Cmd.AddCommand(updateCmd)
	Cmd.AddCommand(rmCmd)

	Cmd.PersistentFlags().StringP("server", "S", "", "The server name to fetch the stack list from")
	Cmd.MarkPersistentFlagRequired("server")

	Cmd.PersistentFlags().StringP("stack", "s", "", "The project/stack to get")
	Cmd.MarkPersistentFlagRequired("stack")

	Cmd.PersistentFlags().StringP("service", "", "", "The service/container to get")
}
