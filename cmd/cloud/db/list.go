package db

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"errors"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/cloud/db"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "List cloud databases",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := db.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))
		options := db.ListOptions{}

		serverNameFlag := cmd.Flag("server")
		if serverNameFlag != nil {
			options.ServerName = serverNameFlag.Value.String()
		}

		databaseFlag := cmd.Flag("database")
		if databaseFlag != nil {
			options.Database = databaseFlag.Value.String()
		}

		hostFlag := cmd.Flag("host")
		if hostFlag != nil {
			options.MySQLHost = hostFlag.Value.String()
		}

		databaseResponse, err := client.List(context.Background(), options)
		if err != nil {
			return err
		}

		format := cmd.Flag("format").Value.String()
		if format == "json" {
			json, err := json.MarshalIndent(databaseResponse.Return.Databases, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(json))
		} else if format == "text" {
			w := new(tabwriter.Writer)
			w.Init(os.Stdout, 0, 4, 4, ' ', 0)
			fmt.Fprintln(w, "Database Id\tDatabase Name\tDatabase Host\tServer Name\tDatabase Container")
			for _, database := range databaseResponse.Return.Databases {
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", database.ID, database.DBName, database.MySQLHost, database.ServerName, database.Container)
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
	listCmd.Flags().StringP("server", "S", "", "The server name")
	listCmd.Flags().StringP("host", "H", "", "The MySQL host")
	listCmd.Flags().StringP("db", "D", "", "The MySQL name")
}
