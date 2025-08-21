package job

import (
	"strconv"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/models"
	"github.com/sitehostnz/terraform-provider-sitehost/sitehost/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var jobWaitCommand = &cobra.Command{
	Use:   "wait",
	Short: "wait for a job to finish",
	RunE: func(cmd *cobra.Command, args []string) error {
		api := api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId"))
		jobType := cmd.Flag("type").Value.String()
		jobId, err := strconv.Atoi(cmd.Flag("job").Value.String())
		if err != nil {
			return err
		}

		return helper.WaitForJob(api, models.Job{ID: jobId, Type: jobType})
	},
}

func init() {
	Cmd.AddCommand(jobWaitCommand)
	Cmd.AddCommand(jobGetCommand)
}
