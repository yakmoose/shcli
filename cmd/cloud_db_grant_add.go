/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"context"
	"github.com/sitehostnz/gosh/pkg/api/cloud/db/grant"
	"github.com/sitehostnz/gosh/pkg/api/job"
	"github.com/sitehostnz/terraform-provider-sitehost/sitehost/helper"
	"strings"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cloudDbUserAddCommand represents the domainAdd command
var cloudDbGrantAddCommand = &cobra.Command{
	Use:   "add",
	Short: "Add new grants for a user",
	RunE: func(cmd *cobra.Command, args []string) error {

		ctx := context.Background()
		api := api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId"))
		userClient := grant.New(api)

		addRequest := grant.AddRequest{
			MySQLHost:  cmd.Flag("host").Value.String(),
			ServerName: cmd.Flag("server").Value.String(),
			Username:   cmd.Flag("user").Value.String(),
			Grants:     strings.Split(cmd.Flag("grants").Value.String(), ","),
			Database:   cmd.Flag("database").Value.String(),
		}

		userAddResponse, err := userClient.Add(ctx, addRequest)
		if nil != err {
			return err
		}

		return helper.WaitForAction(api, job.GetRequest{ID: userAddResponse.Return.Job.ID, Type: userAddResponse.Return.Job.Type})
	},
}

func init() {
	cloudDbGrantCommand.AddCommand(cloudDbGrantAddCommand)

	cloudDbGrantAddCommand.Flags().StringP("grants", "g", "", "The grants to add")
	cloudDbGrantAddCommand.MarkFlagRequired("grants")

}
