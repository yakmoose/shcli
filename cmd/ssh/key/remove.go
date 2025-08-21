package key

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/ssh/key"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command.
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove the specified ssh key",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := key.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))

		keyID := cmd.Flag("keyid").Value.String()

		keyResponse, err := client.Delete(context.Background(), key.DeleteRequest{ID: keyID})
		if err != nil {
			return err
		}

		json, err := json.MarshalIndent(keyResponse, "", "  ")
		fmt.Println(string(json))

		return nil
	},
}

func init() {
	Cmd.AddCommand(rmCmd)

	rmCmd.Flags().StringP("keyid", "", "", "The key id")
	rmCmd.MarkFlagRequired("keyid")
}
