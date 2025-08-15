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

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a database user",
	RunE: func(cmd *cobra.Command, args []string) error {

		client := user.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))

		userResponse, err := client.Get(
			context.Background(),
			user.GetRequest{
				ServerName: cmd.Flag("server").Value.String(),
				Username:   cmd.Flag("user").Value.String(),
				MySQLHost:  cmd.Flag("host").Value.String(),
			})

		if err != nil {
			return err
		}

		json, err := json.MarshalIndent(userResponse, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(json))

		return nil
	},
}

func init() {
	getCmd.Flags().StringP("server", "S", "", "The server where the user exists")
	getCmd.MarkFlagRequired("server")

	getCmd.Flags().StringP("user", "U", "", "The MySQL username")
	getCmd.MarkFlagRequired("user")

	getCmd.Flags().StringP("host", "H", "", "The MySQL host")
	getCmd.MarkFlagRequired("host")

}
