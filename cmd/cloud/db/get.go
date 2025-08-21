package db

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/cloud/db"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the database",
	RunE: func(cmd *cobra.Command, args []string) error {

		client := db.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))

		databaseResponse, err := client.Get(
			context.Background(),
			db.GetRequest{
				Database:   cmd.Flag("db").Value.String(),
				ServerName: cmd.Flag("server").Value.String(),
				MySQLHost:  cmd.Flag("host").Value.String(),
			})

		if err != nil {
			return err
		}

		json, err := json.MarshalIndent(databaseResponse, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(json))

		return nil
	},
}

func init() {

	getCmd.Flags().StringP("server", "S", "", "The server name")
	getCmd.MarkFlagRequired("server")

	getCmd.Flags().StringP("host", "H", "", "The database host")
	getCmd.MarkFlagRequired("host")

	getCmd.Flags().StringP("db", "d", "", "The database name")
	getCmd.MarkFlagRequired("db")

}
