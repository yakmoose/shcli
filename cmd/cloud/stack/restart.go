package stack

import (
	"context"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/cloud/stack"
	"github.com/sitehostnz/terraform-provider-sitehost/sitehost/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// restartCmd represents the list command
var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restarts a stack",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := viper.GetString("apiKey")
		clientId := viper.GetString("clientId")
		api := api.NewClient(apiKey, clientId)
		client := stack.New(api)

		serverName := cmd.Flag("server").Value.String()
		stackName := cmd.Flag("stack").Value.String()
		response, err := client.Restart(context.Background(), stack.StopStartRestartRequest{ServerName: serverName, Name: stackName})
		if err != nil {
			return err
		}

		return helper.WaitForJob(api, response.Return.Job)

	},
}

func init() {
	restartCmd.Flags().StringP("server", "S", "", "The server name to fetch the stack list from")
	restartCmd.MarkFlagRequired("server")

	restartCmd.Flags().StringP("stack", "s", "", "The stack to restart")
	restartCmd.MarkFlagRequired("stack")
}
