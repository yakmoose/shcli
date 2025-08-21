package grant

import (
	"context"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/cloud/db/grant"
	"github.com/sitehostnz/terraform-provider-sitehost/sitehost/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rmCmd drops/deletes the grants for a user
var rmCmd = &cobra.Command{
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

		response, err := userClient.Delete(ctx, deleteRequest)
		if nil != err {
			return err
		}

		return helper.WaitForJob(api, response.Return.Job)

	},
}
