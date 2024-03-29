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
	"github.com/sitehostnz/gosh/pkg/api/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listServersCmd = &cobra.Command{
	Use:   "list",
	Short: "List servers",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := server.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))

		serversResponse, err := client.List(context.Background())
		if err != nil {
			return err
		}

		format := cmd.Flag("format").Value.String()
		if format == "json" {
			json, err := json.MarshalIndent(serversResponse.Return.Servers, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(json))
		} else if format == "text" {
			w := new(tabwriter.Writer)
			w.Init(os.Stdout, 0, 4, 4, ' ', 0)
			fmt.Fprintln(w, "Server Name\tServer Label\tProduct Type\tServer Cores\tServer Ram\tServer Disk")
			for _, server := range serversResponse.Return.Servers {
				fmt.Fprintf(w, "%s\t%s\t%s\t%d\t%s\t%d\n", server.Name, server.Label, server.ProductType, server.Cores, server.RAM, int64(server.Disk))
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
	serverCmd.AddCommand(listServersCmd)
}
