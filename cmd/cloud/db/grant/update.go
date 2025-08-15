package grant

import (
	"context"
	"strings"

	"github.com/sitehostnz/gosh/pkg/api/cloud/db/grant"
	"github.com/sitehostnz/terraform-provider-sitehost/sitehost/helper"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// updateCmd updates a users grants
var updateCmd = &cobra.Command{
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

		response, err := userClient.Update(ctx, updateRequest)
		if nil != err {
			return err
		}

		return helper.WaitForJob(api, response.Return.Job)

	},
}

func init() {
	updateCmd.Flags().StringP("grants", "g", "", "The grants to update")
	updateCmd.MarkFlagRequired("grants")
}
