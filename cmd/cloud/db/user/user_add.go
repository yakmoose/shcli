package user

import (
	"context"

	"github.com/sitehostnz/gosh/pkg/api/cloud/db/user"
	"github.com/sitehostnz/terraform-provider-sitehost/sitehost/helper"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the domainAdd command
var addCmd = &cobra.Command{
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

		response, err := userClient.Add(ctx, addRequest)
		if nil != err {
			return err
		}

		return helper.WaitForJob(api, response.Return.Job)
	},
}

func init() {
	addCmd.Flags().StringP("server", "S", "", "The server name")
	addCmd.MarkFlagRequired("server")

	addCmd.Flags().StringP("host", "H", "", "The database host")
	addCmd.MarkFlagRequired("host")

	addCmd.Flags().StringP("user", "U", "", "The database user")
	addCmd.MarkFlagRequired("user")

	addCmd.Flags().StringP("password", "p", "", "The database user password")
	addCmd.MarkFlagRequired("password")
}
