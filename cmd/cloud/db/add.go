package db

import (
	"context"
	"errors"

	"github.com/sitehostnz/gosh/pkg/api/cloud/db"
	"github.com/sitehostnz/gosh/pkg/api/server"
	"github.com/sitehostnz/terraform-provider-sitehost/sitehost/helper"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the domainAdd command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new database",
	RunE: func(cmd *cobra.Command, args []string) error {

		ctx := context.Background()
		api := api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId"))
		dbClient := db.New(api)
		serverClient := server.New(api)

		database := cmd.Flag("db").Value.String()
		serverName := cmd.Flag("server").Value.String()
		host := cmd.Flag("host").Value.String()
		container := cmd.Flag("container").Value.String()

		// 1. is the server a stack server? best check
		stackServer, err := serverClient.Get(context.Background(), server.GetRequest{ServerName: serverName})
		if err != nil {
			return err
		}

		if stackServer.Server.ProductType != "CLDCON" {
			return errors.New("server is not a cloud container server")
		}

		response, err := dbClient.Add(ctx, db.AddRequest{Database: database, MySQLHost: host, ServerName: serverName, Container: container})
		if err != nil {
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

	addCmd.Flags().StringP("db", "d", "", "The database name")
	addCmd.MarkFlagRequired("db")

	addCmd.Flags().StringP("container", "c", "", "The database backup container")
	addCmd.MarkFlagRequired("container")

}
