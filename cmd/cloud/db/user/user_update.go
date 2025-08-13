package user

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/cloud/db/user"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// updateCmd updates the cloud db backup container
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the database users password",
	RunE: func(cmd *cobra.Command, args []string) error {

		ctx := context.Background()
		client := user.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))

		serverName := cmd.Flag("server").Value.String()
		host := cmd.Flag("host").Value.String()
		userName := cmd.Flag("user").Value.String()
		password := cmd.Flag("password").Value.String()

		response, err := client.Update(
			ctx,
			user.UpdateRequest{
				MySQLHost:  host,
				ServerName: serverName,
				Username:   userName,
				Password:   password,
			})
		if err != nil {
			return err
		}

		if err != nil {
			return err
		}

		json, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(json))

		return nil
	},
}

func init() {

	updateCmd.Flags().StringP("server", "S", "", "The server name")
	updateCmd.MarkFlagRequired("server")

	updateCmd.Flags().StringP("host", "H", "", "The database host")
	updateCmd.MarkFlagRequired("host")

	updateCmd.Flags().StringP("user", "U", "", "The database user")
	updateCmd.MarkFlagRequired("user")

	updateCmd.Flags().StringP("password", "p", "", "The database user password")
	updateCmd.MarkFlagRequired("password")
}
