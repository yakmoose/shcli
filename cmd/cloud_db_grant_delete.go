/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"context"
	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/cloud/db/grant"
	"github.com/sitehostnz/gosh/pkg/api/job"
	"github.com/sitehostnz/terraform-provider-sitehost/sitehost/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cloudDbGrantDeleteCommand drops/deletes the grants for a user
var cloudDbGrantDeleteCommand = &cobra.Command{
	Use:   "drop",
	Short: "drop the user grants",
	RunE: func(cmd *cobra.Command, args []string) error {

		ctx := context.Background()
		api := api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId"))
		userClient := grant.New(api)

		deleteRequest := grant.DeleteRequest{
			MySQLHost:  cmd.Flag("host").Value.String(),
			ServerName: cmd.Flag("server").Value.String(),
			Username:   cmd.Flag("user").Value.String(),
			Database:   cmd.Flag("database").Value.String(),
		}

		deleteResponse, err := userClient.Delete(ctx, deleteRequest)
		if nil != err {
			return err
		}

		return helper.WaitForAction(api, job.GetRequest{ID: deleteResponse.Return.Job.ID, Type: deleteResponse.Return.Job.Type})

	},
}

func init() {
	cloudDbGrantCommand.AddCommand(cloudDbGrantDeleteCommand)
}
