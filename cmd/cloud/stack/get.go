package stack

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/cloud/stack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the stack",
	RunE: func(cmd *cobra.Command, args []string) error {

		client := stack.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))

		serverName := cmd.Flag("server").Value.String()
		stackName := cmd.Flag("stack").Value.String()

		stack, err := client.Get(context.Background(), stack.GetRequest{ServerName: serverName, Name: stackName})
		if err != nil {
			return err
		}

		json, err := json.MarshalIndent(stack, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(json))

		return nil
	},
}

func init() {
	getCmd.Flags().StringP("server", "S", "", "The server name to fetch the stack list from")
	getCmd.MarkFlagRequired("server")

	getCmd.Flags().StringP("stack", "s", "", "The stack to get")
	getCmd.MarkFlagRequired("stack")
}
