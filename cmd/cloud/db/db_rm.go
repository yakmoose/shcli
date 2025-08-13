package db

import (
	"context"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/cloud/db"
	"github.com/sitehostnz/terraform-provider-sitehost/sitehost/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rmCmd represents the domainAdd command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "delete a database",
	RunE: func(cmd *cobra.Command, args []string) error {

		ctx := context.Background()
		api := api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId"))
		client := db.New(api)

		database := cmd.Flag("db").Value.String()
		serverName := cmd.Flag("server").Value.String()
		host := cmd.Flag("host").Value.String()

		response, err := client.Delete(ctx, db.DeleteRequest{Database: database, MySQLHost: host, ServerName: serverName})
		if err != nil {
			return err
		}

		return helper.WaitForJob(api, response.Return.Job)
	},
}

func init() {

	rmCmd.Flags().StringP("server", "S", "", "The server name")
	rmCmd.MarkFlagRequired("server")

	rmCmd.Flags().StringP("host", "H", "", "The database host")
	rmCmd.MarkFlagRequired("host")

	rmCmd.Flags().StringP("db", "d", "", "The database name")
	rmCmd.MarkFlagRequired("db")

}
