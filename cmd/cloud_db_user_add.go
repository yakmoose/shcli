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

// cloudDbUserAddCommand represents the domainAdd command
var cloudDbUserAddCommand = &cobra.Command{
	Use:   "add",
	Short: "Add a new database user",
	RunE: func(cmd *cobra.Command, args []string) error {

		ctx := context.Background()
		api := api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId"))
		userClient := user.New(api)

		//database := cmd.Flag("db").Value.String()
		//serverName := cmd.Flag("server").Value.String()
		//host := cmd.Flag("host").Value.String()

		addRequest := user.AddRequest{
			MySQLHost:  cmd.Flag("host").Value.String(),
			ServerName: cmd.Flag("server").Value.String(),
			Username:   cmd.Flag("user").Value.String(),
			Password:   cmd.Flag("password").Value.String(),
		}

		userAddResponse, err := userClient.Add(ctx, addRequest)
		if nil != err {
			return err
		}

		return helper.WaitForAction(api, job.GetRequest{JobID: userAddResponse.Return.JobID, Type: job.SchedulerType})
	},
}

func init() {
	cloudDbUserCommand.AddCommand(cloudDbUserAddCommand)

	cloudDbUserAddCommand.Flags().StringP("server", "S", "", "The server name")
	cloudDbUserAddCommand.MarkFlagRequired("server")

	cloudDbUserAddCommand.Flags().StringP("host", "H", "", "The database host")
	cloudDbUserAddCommand.MarkFlagRequired("host")

	cloudDbUserAddCommand.Flags().StringP("user", "U", "", "The database user")
	cloudDbUserAddCommand.MarkFlagRequired("user")

	cloudDbUserAddCommand.Flags().StringP("password", "p", "", "The database user password")
	cloudDbUserAddCommand.MarkFlagRequired("password")

}
