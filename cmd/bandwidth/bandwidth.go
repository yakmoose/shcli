package bandwidth

import (
	"github.com/spf13/cobra"
)

// Cmd root bandwidth comand
var Cmd = &cobra.Command{
	Use:   "bandwidth",
	Short: "Commands viewing and managing bandwidth data",
}
