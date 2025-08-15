package environment

import (
	"context"
	"encoding/json"
	"io"
	"os"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/cloud/stack/environment"
	"github.com/sitehostnz/gosh/pkg/models"
	"github.com/sitehostnz/terraform-provider-sitehost/sitehost/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rmCmd represents the get command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Deletes the supplied variables from the stack environment",
	RunE: func(cmd *cobra.Command, args []string) error {
		api := api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId"))
		client := environment.New(api)

		serverName := cmd.Flag("server").Value.String()
		stackName := cmd.Flag("stack").Value.String()
		serviceName := cmd.Flag("service").Value.String()
		fileName := cmd.Flag("file").Value.String()

		if "" == serviceName {
			serviceName = stackName
		}

		var fd *os.File
		var err error
		if len(fileName) > 0 {
			fd, err = os.Open(fileName)
			if nil != err {
				return err
			}
		} else {
			fd = os.Stdin
		}

		//

		data, err := io.ReadAll(fd)
		if err != nil {
			return err
		}

		var settings []models.EnvironmentVariable

		err = json.Unmarshal(data, &settings)
		if err != nil {
			return err
		}

		response, err := client.Delete(context.Background(), environment.DeleteRequest{
			ServerName:           serverName,
			Project:              stackName,
			Service:              serviceName,
			EnvironmentVariables: settings,
		})

		if err != nil {
			return err
		}

		return helper.WaitForJob(api, response.Return.Job)

	},
}

func init() {
	rmCmd.Flags().StringP("file", "F", "", "The settings json file to use")
}
