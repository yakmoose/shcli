package user

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/cloud/ssh/user"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the user",
	RunE: func(cmd *cobra.Command, args []string) error {

		client := user.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))

		databaseResponse, err := client.Get(
			context.Background(),
			user.GetRequest{
				ServerName: cmd.Flag("server").Value.String(),
				Username:   cmd.Flag("user").Value.String(),
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
	Cmd.AddCommand(getCmd)

	getCmd.Flags().StringP("server", "S", "", "The server name")
	getCmd.MarkFlagRequired("server")

	getCmd.Flags().StringP("user", "u", "", "The username")
	getCmd.MarkFlagRequired("user")

}
