package cloud

import (
	"shcli/cmd/cloud/db"
	"shcli/cmd/cloud/image"
	"shcli/cmd/cloud/ssh"
	"shcli/cmd/cloud/stack"
	"shcli/cmd/cloud/stack/server"

	"github.com/spf13/cobra"
)

// Cmd represents the cloud command
var Cmd = &cobra.Command{
	Use:   "cloud",
	Short: "Commands for manipulating with Sitehost cloud/container stacks",
}

func init() {
	Cmd.AddCommand(stack.Cmd)
	Cmd.AddCommand(server.Cmd)
	Cmd.AddCommand(image.Cmd)
	Cmd.AddCommand(db.Cmd)
	Cmd.AddCommand(ssh.Cmd)
}
