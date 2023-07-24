/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"errors"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/cloud/db/user"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cloudDbListCmd represents the list command
var cloudDbUserListCmd = &cobra.Command{
	Use:   "list",
	Short: "List cloud database users",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := user.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))
		options := user.ListOptions{}

		serverNameFlag := cmd.Flag("server")
		if serverNameFlag != nil {
			options.ServerName = serverNameFlag.Value.String()
		}

		userNameFlag := cmd.Flag("user")
		if userNameFlag != nil {
			options.Username = userNameFlag.Value.String()
		}

		mysqlHostFlag := cmd.Flag("host")
		if mysqlHostFlag != nil {
			options.MySQLHost = mysqlHostFlag.Value.String()
		}

		databaseUserResponse, err := client.List(context.Background(), options)
		if err != nil {
			return err
		}

		format := cmd.Flag("format").Value.String()
		if format == "json" {
			json, err := json.MarshalIndent(databaseUserResponse.Return.Users, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(json))
		} else if format == "text" {
			w := new(tabwriter.Writer)
			w.Init(os.Stdout, 0, 4, 4, ' ', 0)
			fmt.Fprintln(w, "Username\tDatabase Host\tServer Name")
			for _, user := range databaseUserResponse.Return.Users {
				fmt.Fprintf(w, "%s\t%s\t%s\n", user.Username, user.MysqlHost, user.ServerName)
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
	cloudDbUserCommand.AddCommand(cloudDbUserListCmd)
	cloudDbUserListCmd.Flags().StringP("user", "U", "", "The MySQL username")
	cloudDbUserListCmd.Flags().StringP("server", "S", "", "The MySQL server")
	cloudDbUserListCmd.Flags().StringP("host", "H", "", "The MySQL host")

}
