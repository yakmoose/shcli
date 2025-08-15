package user

import (
	"context"

	"github.com/sitehostnz/gosh/pkg/api/cloud/db/user"
	"github.com/sitehostnz/terraform-provider-sitehost/sitehost/helper"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// dropCmd represents the domainAdd command
var dropCmd = &cobra.Command{
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

		return helper.WaitForJob(api, response.Return.Job)
	},
}

func init() {
	dropCmd.Flags().StringP("server", "S", "", "The server name")
	dropCmd.MarkFlagRequired("server")

	dropCmd.Flags().StringP("host", "H", "", "The database host")
	dropCmd.MarkFlagRequired("host")

	dropCmd.Flags().StringP("user", "U", "", "The database user")
	dropCmd.MarkFlagRequired("user")
}
