package grant

import (
	"github.com/spf13/cobra"
)

// Cmd represents the stacks command
var Cmd = &cobra.Command{
	Use:   "grant",
	Short: "Manage cloud database grants",
}

func init() {

	Cmd.AddCommand(addCmd)
	Cmd.AddCommand(rmCmd)
	Cmd.AddCommand(updateCmd)

	Cmd.PersistentFlags().StringP("server", "S", "", "The server name")
	Cmd.MarkPersistentFlagRequired("server")

	Cmd.PersistentFlags().StringP("host", "H", "", "The database host")
	Cmd.MarkPersistentFlagRequired("host")

	Cmd.PersistentFlags().StringP("user", "U", "", "The database user")
	Cmd.MarkPersistentFlagRequired("user")

	Cmd.PersistentFlags().StringP("database", "d", "", "The database to add grants to")
	Cmd.MarkPersistentFlagRequired("database")

}
