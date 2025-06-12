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

// cloudDbGrantUpdateCommand updates a users grants
var cloudDbGrantUpdateCommand = &cobra.Command{
	Use:   "update",
	Short: "Update new grants for a user",
	RunE: func(cmd *cobra.Command, args []string) error {

		ctx := context.Background()
		api := api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId"))
		userClient := grant.New(api)

		updateRequest := grant.UpdateRequest{
			MySQLHost:  cmd.Flag("host").Value.String(),
			ServerName: cmd.Flag("server").Value.String(),
			Username:   cmd.Flag("user").Value.String(),
			Database:   cmd.Flag("database").Value.String(),
			Grants:     strings.Split(cmd.Flag("grants").Value.String(), ","),
		}

		updateResponse, err := userClient.Update(ctx, updateRequest)
		if nil != err {
			return err
		}

		return helper.WaitForAction(api, job.GetRequest{ID: updateResponse.Return.Job.ID, Type: updateResponse.Return.Job.Type})

	},
}

func init() {
	cloudDbGrantCommand.AddCommand(cloudDbGrantUpdateCommand)

	cloudDbGrantUpdateCommand.Flags().StringP("grants", "g", "", "The grants to update")
	cloudDbGrantUpdateCommand.MarkFlagRequired("grants")
}
