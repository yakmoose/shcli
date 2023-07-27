package cmd

import "github.com/spf13/cobra"

var bandwidthCommand = &cobra.Command{
	Use:   "bandwidth",
	Short: "Commands viewing and managing bandwidth data",
}

func init() {
	rootCmd.AddCommand(bandwidthCommand)
}
