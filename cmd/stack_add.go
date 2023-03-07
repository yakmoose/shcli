/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/cloud/stack"
	"github.com/sitehostnz/gosh/pkg/api/server"
	"github.com/sitehostnz/gosh/pkg/models"
	"github.com/sitehostnz/terraform-provider-sitehost/sitehost/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"os"
)

// getCmd represents the get command
var stackAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a stack",
	RunE: func(cmd *cobra.Command, args []string) error {

		serverName := cmd.Flag("server").Value.String()
		stackLabel := cmd.Flag("label").Value.String()
		stackName := cmd.Flag("name").Value.String()
		dockerFileName := cmd.Flag("compose-file").Value.String()
		environmentFileName := cmd.Flag("environment-file").Value.String()
		enableSSL, _ := cmd.Flags().GetBool("enableSSL")

		// read the docker file.
		var fd *os.File
		var err error
		if len(dockerFileName) > 0 {
			fd, err = os.Open(dockerFileName)
			if nil != err {
				return err
			}
		} else {
			fd = os.Stdin
		}

		composeFile, err := io.ReadAll(fd)

		var environmentVariables []models.EnvironmentVariable
		if len(environmentFileName) > 0 {
			fd, err = os.Open(environmentFileName)
			if nil != err {
				return err
			}

			environmentFile, err := io.ReadAll(fd)

			err = json.Unmarshal(environmentFile, &environmentVariables)
			if err != nil {
				return err
			}

		} else {
			environmentVariables = []models.EnvironmentVariable{}
		}

		client := api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId"))
		serverClient := server.New(client)
		stackClient := stack.New(client)

		// 1. is the server a stack server? best check
		stackServer, err := serverClient.Get(context.Background(), server.GetRequest{ServerName: serverName})
		if err != nil {
			return err
		}

		if stackServer.Server.ProductType != "CLDCON" {
			return errors.New("server is not a cloud container server")
		}

		stackResponse, err := stackClient.Add(
			context.Background(),
			stack.AddRequest{
				ServerName: serverName,
				Name:       stackName,
				Label:      stackLabel,
				EnableSSL: func(v bool) int {
					if v {
						return 1
					} else {
						return 0
					}
				}(enableSSL),
				DockerCompose:        string(composeFile),
				EnvironmentVariables: environmentVariables,
			},
		)

		if err != nil {
			return err
		}

		return helper.WaitForAction(client, stackResponse.Return.JobID)

	},
}

func init() {
	stackCommand.AddCommand(stackAddCmd)

	stackAddCmd.Flags().StringP("server", "S", "", "The server add the stack to")
	stackAddCmd.MarkFlagRequired("server")

	stackAddCmd.Flags().StringP("label", "l", "", "The stack label")
	stackAddCmd.MarkFlagRequired("label")

	stackAddCmd.Flags().StringP("name", "n", "", "The stack name, should be generated by a call to generatename")
	stackAddCmd.MarkFlagRequired("name")

	stackAddCmd.Flags().StringP("compose-file", "c", "", "The docker compose file")
	stackAddCmd.MarkFlagRequired("compose-file")

	stackAddCmd.Flags().StringP("environment-file", "e", "", "The environment")

	stackAddCmd.Flags().BoolP("enableSSL", "", false, "Enable ssl")

}
