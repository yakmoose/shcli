/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/ssh/key"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command.
var listKeysCmd = &cobra.Command{
	Use:   "list",
	Short: "List ssh keys",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := key.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))

		keysResponse, err := client.List(context.Background())
		if err != nil {
			return err
		}

		format := cmd.Flag("format").Value.String()

		if format == "json" {
			json, err := json.MarshalIndent(keysResponse.Return.SSHKeys, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(json))
		} else if format == "text" {
			w := new(tabwriter.Writer)
			w.Init(os.Stdout, 0, 4, 4, ' ', 0)
			fmt.Fprintln(w, "Id\tLabel\tKey")
			for _, key := range keysResponse.Return.SSHKeys {
				fmt.Fprintf(w, "%s\t%s\t%s\n", key.ID, key.Label, key.Content)
			}

			fmt.Fprintln(w)
		} else {
			/// how to error out.
			return errors.New("unsupported format, please choose text or json")
		}

		return nil
	},
}

func init() {
	sshKeysCmd.AddCommand(listKeysCmd)
}
