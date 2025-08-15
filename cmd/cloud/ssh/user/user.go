package user

import (
	"github.com/spf13/cobra"
)

// Cmd represents the user command
var Cmd = &cobra.Command{
	Use:   "user",
	Short: "Manage cloud ssh/sftp users",
}
