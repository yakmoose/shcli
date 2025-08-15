package key

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/ssh/key"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"context"
	"os"
)

// addCmd represents the add command.
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a ssh key",
	RunE: func(cmd *cobra.Command, args []string) error {

		label := cmd.Flag("label").Value.String()
		keyFileName := cmd.Flag("key").Value.String()

		// read the docker file.
		var fd *os.File
		var err error
		if len(keyFileName) > 0 {
			fd, err = os.Open(keyFileName)
			if nil != err {
				return err
			}
		} else {
			fd = os.Stdin
		}

		keyFile, err := io.ReadAll(fd)

		client := api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId"))
		keyClient := key.New(client)

		// 1. is the server a stack server? best check
		keyResponse, err := keyClient.Create(context.Background(), key.CreateRequest{Content: string(keyFile), Label: label})
		if err != nil {
			return err
		}

		json, err := json.MarshalIndent(keyResponse.Return, "", "  ")
		fmt.Println(string(json))

		return nil
	},
}

func init() {
	Cmd.AddCommand(addCmd)

	addCmd.Flags().StringP("label", "l", "", "The key label")
	addCmd.MarkFlagRequired("label")

	addCmd.Flags().StringP("key", "k", "", "The key file")

}
