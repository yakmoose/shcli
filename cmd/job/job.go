package job

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "job",
	Short: "Commands getting information about jobs",
}

func init() {
	Cmd.PersistentFlags().StringP("job", "", "", "The job id")
	Cmd.MarkFlagRequired("job")

	Cmd.PersistentFlags().StringP("type", "", "", "The job type")
	Cmd.MarkFlagRequired("type")
}
