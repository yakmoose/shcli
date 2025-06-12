package cmd

import (
	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/job"
	"github.com/sitehostnz/terraform-provider-sitehost/sitehost/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strconv"
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

		return helper.WaitForAction(api, job.GetRequest{ID: jobId, Type: jobType})
	},
}

func init() {
	jobCommand.AddCommand(jobWaitCommand)
	jobCommand.AddCommand(jobGetCommand)
}
