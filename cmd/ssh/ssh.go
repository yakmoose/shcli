package ssh

import (
	"shcli/cmd/cloud/ssh"
	"shcli/cmd/ssh/key"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "ssh",
	Short: "Manage ssh",
}

func init() {
	ssh.Cmd.AddCommand(key.Cmd)
}
