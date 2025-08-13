package environment

import (
	"github.com/spf13/cobra"
)

// EnvironmentCmd represents the domain command
var EnvironmentCmd = &cobra.Command{
	Use:   "env",
	Short: "Commands for managing stack environment variables",
}

func init() {

	EnvironmentCmd.PersistentFlags().StringP("server", "S", "", "The server name to fetch the stack list from")
	EnvironmentCmd.MarkPersistentFlagRequired("server")

	EnvironmentCmd.PersistentFlags().StringP("stack", "s", "", "The project/stack to get")
	EnvironmentCmd.MarkPersistentFlagRequired("stack")

	EnvironmentCmd.PersistentFlags().StringP("service", "", "", "The service/container to get")
}
