package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/cloud/ssh/user"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
	"text/tabwriter"
)

// cloudSSHUserList represents the get command
var cloudSSHUserList = &cobra.Command{
	Use:   "list",
	Short: "List the users",
	RunE: func(cmd *cobra.Command, args []string) error {

		client := user.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))

		response, err := client.List(
			context.Background(),
			user.ListOptions{
				ServerName: cmd.Flag("server").Value.String(),
				Username:   cmd.Flag("username").Value.String(),
			})

		if err != nil {
			return err
		}

		format := cmd.Flag("format").Value.String()
		if format == "json" {
			json, err := json.MarshalIndent(response.Return.Users, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(json))
		} else if format == "text" {
			w := new(tabwriter.Writer)
			w.Init(os.Stdout, 0, 4, 4, ' ', 0)
			fmt.Fprintln(w, "User Name\tServer Name\tContainers")
			for _, user := range response.Return.Users {
				fmt.Fprintf(w, "%s\t%s\t%s\n", user.Username, user.ServerName, strings.Join(user.Containers, ", "))
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
	cloudSSHUserCmd.AddCommand(cloudSSHUserList)
	cloudSSHUserList.Flags().StringP("server", "S", "", "The server name")
	cloudSSHUserList.Flags().StringP("username", "U", "", "The username")
}
