package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/job"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strconv"
)

var jobGetCommand = &cobra.Command{
	Use:   "get",
	Short: "get details about the specified job",
	RunE: func(cmd *cobra.Command, args []string) error {
		api := api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId"))
		jobType := cmd.Flag("type").Value.String()
		jobId, err := strconv.Atoi(cmd.Flag("job").Value.String())
		if err != nil {
			return err
		}

		client := job.New(api)
		response, err := client.Get(context.Background(), job.GetRequest{ID: jobId, Type: jobType})
		if err != nil {
			return err
		}

		json, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			return err
		}
		
		fmt.Println(string(json))

		return nil
	},
}

func init() {
	jobCommand.AddCommand(jobGetCommand)
}
