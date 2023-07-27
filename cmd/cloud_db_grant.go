/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// cloudDbGrantCommand represents the stacks command
var cloudDbGrantCommand = &cobra.Command{
	Use:   "grant",
	Short: "Manage cloud database grants",
}

func init() {
	cloudDbCommand.AddCommand(cloudDbGrantCommand)

	cloudDbGrantCommand.PersistentFlags().StringP("server", "S", "", "The server name")
	cloudDbGrantCommand.MarkPersistentFlagRequired("server")

	cloudDbGrantCommand.PersistentFlags().StringP("host", "H", "", "The database host")
	cloudDbGrantCommand.MarkPersistentFlagRequired("host")

	cloudDbGrantCommand.PersistentFlags().StringP("user", "U", "", "The database user")
	cloudDbGrantCommand.MarkPersistentFlagRequired("user")

	cloudDbGrantCommand.PersistentFlags().StringP("database", "d", "", "The database to add grants to")
	cloudDbGrantCommand.MarkPersistentFlagRequired("database")

}
