/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// cloudDbCommand represents the stacks command
var cloudDbCommand = &cobra.Command{
	Use:   "db",
	Short: "Manage cloud databases",
}

func init() {
	cloudCommand.AddCommand(cloudDbCommand)
}
