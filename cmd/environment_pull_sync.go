/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"context"
	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/cloud/stack/environment"
	"github.com/sitehostnz/gosh/pkg/api/job"
	"github.com/sitehostnz/gosh/pkg/models"
	"github.com/sitehostnz/terraform-provider-sitehost/sitehost/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yakmoose/envop/service"
	"strings"
)

// environmentPullCmd represents the get command
var environmentPullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull environment/item from 1Password and push it to Sitehost",
	RunE: func(cmd *cobra.Command, args []string) error {
		api := api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId"))
		client := environment.New(api)

		serverName := cmd.Flag("server").Value.String()
		stackName := cmd.Flag("stack").Value.String()
		serviceName := cmd.Flag("service").Value.String()
		vaultName := cmd.Flag("vault").Value.String()
		itemName := cmd.Flag("item").Value.String()

		if "" == serviceName {
			serviceName = stackName
		}

		token, err := cmd.Flags().GetString("service-account")
		if err != nil {
			return err
		}

		op, err := service.NewClientFromToken(token)
		if err != nil {
			return err
		}

		item, err := service.Get1PasswordItem(op, vaultName, itemName)
		if err != nil {
			return err
		}

		var settings []models.EnvironmentVariable

		for _, v := range item.Fields {
			if v.Value == "" {
				continue
			}
			settings = append(settings, models.EnvironmentVariable{
				Name:    strings.ToUpper(v.Title),
				Content: v.Value,
			})
		}

		response, err := client.Update(context.Background(), environment.UpdateRequest{
			ServerName:           serverName,
			Project:              stackName,
			Service:              serviceName,
			EnvironmentVariables: settings,
		})

		if err != nil {
			return err
		}

		return helper.WaitForAction(api, job.GetRequest{ID: response.Return.Job.ID, Type: response.Return.Job.Type})

	},
}

func init() {
	environmentOpCmd.AddCommand(environmentPullCmd)
}
