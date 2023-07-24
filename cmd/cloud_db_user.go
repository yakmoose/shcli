/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// cloudDbUserCommand represents the db user commands
var cloudDbUserCommand = &cobra.Command{
	Use:   "user",
	Short: "Manage cloud database users",
}

func init() {
	cloudDbCommand.AddCommand(cloudDbUserCommand)
}
