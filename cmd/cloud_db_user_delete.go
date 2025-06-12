/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"context"
	"github.com/sitehostnz/gosh/pkg/api/cloud/db/user"
	"github.com/sitehostnz/gosh/pkg/api/job"
	"github.com/sitehostnz/terraform-provider-sitehost/sitehost/helper"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cloudDbUserDeleteCommand represents the domainAdd command
var cloudDbUserDeleteCommand = &cobra.Command{
	Use:   "drop",
	Short: "Remove a database user",
	RunE: func(cmd *cobra.Command, args []string) error {

		ctx := context.Background()
		api := api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId"))
		userClient := user.New(api)

		deleteRequest := user.DeleteRequest{
			MySQLHost:  cmd.Flag("host").Value.String(),
			ServerName: cmd.Flag("server").Value.String(),
			Username:   cmd.Flag("user").Value.String(),
		}

		response, err := userClient.Delete(ctx, deleteRequest)
		if nil != err {
			return err
		}

		return helper.WaitForAction(api, job.GetRequest{ID: response.Return.Job.ID, Type: response.Return.Job.Type})
	},
}

func init() {
	cloudDbUserCommand.AddCommand(cloudDbUserDeleteCommand)

	cloudDbUserDeleteCommand.Flags().StringP("server", "S", "", "The server name")
	cloudDbUserDeleteCommand.MarkFlagRequired("server")

	cloudDbUserDeleteCommand.Flags().StringP("host", "H", "", "The database host")
	cloudDbUserDeleteCommand.MarkFlagRequired("host")

	cloudDbUserDeleteCommand.Flags().StringP("user", "U", "", "The database user")
	cloudDbUserDeleteCommand.MarkFlagRequired("user")
}
