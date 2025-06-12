package cmd

import "github.com/spf13/cobra"

var jobCommand = &cobra.Command{
	Use:   "job",
	Short: "Commands getting information about jobs",
}

func init() {
	rootCmd.AddCommand(jobCommand)

	jobCommand.PersistentFlags().StringP("job", "", "", "The job id")
	jobCommand.MarkFlagRequired("job")

	jobCommand.PersistentFlags().StringP("type", "", "", "The job type")
	jobCommand.MarkFlagRequired("type")
}
