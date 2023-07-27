/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/1Password/connect-sdk-go/connect"
	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/cloud/stack/environment"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yakmoose/envop/service"
)

// environmentPushCmd represents the get command
var environmentPushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push environment to 1Password from Sitehost, creating a new password item",
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

		opClient, err := connect.NewClientFromEnvironment()
		if err != nil {
			return err
		}

		response, err := client.Get(context.Background(), environment.GetRequest{
			ServerName: serverName,
			Project:    stackName,
			Service:    serviceName,
		})

		if err != nil {
			return err
		}

		environment := map[string]string{}
		for _, v := range response.EnvironmentVariables {
			environment[v.Name] = v.Content
		}

		item, err := service.Create1PasswordItem(opClient, vaultName, itemName, environment)

		if err != nil {
			return err
		}

		fmt.Printf("item created: %s (%s)\n", item.Title, item.ID)

		return nil
	},
}

func init() {
	environmentOpCmd.AddCommand(environmentPushCmd)

	environmentPullCmd.Flags().StringP("vault", "V", "", "The 1password vault")
	environmentPullCmd.MarkFlagRequired("vault")

	environmentPullCmd.Flags().String("item", "i", "The name of the item to save")
	environmentPullCmd.MarkFlagRequired("item")
}
