package stack

import (
	"shcli/cmd/cloud/stack/environment"

	"github.com/spf13/cobra"
)

// Cmd represents the stacks command
var Cmd = &cobra.Command{
	Use:   "stack",
	Short: "Manage cloud stacks",
}

func init() {
	Cmd.AddCommand(environment.EnvironmentCmd)

	Cmd.AddCommand(addCmd)
	Cmd.AddCommand(findCmd)
	Cmd.AddCommand(generateNameCmd)
	Cmd.AddCommand(getCmd)
	Cmd.AddCommand(getDockerFileCmd)
	Cmd.AddCommand(listCmd)
	Cmd.AddCommand(listVirtualHostsCmd)
	Cmd.AddCommand(restartCmd)
	Cmd.AddCommand(startCmd)
	Cmd.AddCommand(stopCmd)

}
