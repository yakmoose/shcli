package srs

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/srs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"context"
	"encoding/json"
	"errors"
)

var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all registered domain names",
	RunE: func(cmd *cobra.Command, args []string) error {

		client := srs.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))

		response, err := client.ListDomains(context.Background())
		if err != nil {
			return err
		}

		format := cmd.Flag("format").Value.String()
		if format == "json" {
			json, err := json.MarshalIndent(response.Return.Domains, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(json))
		} else if format == "text" {
			w := new(tabwriter.Writer)
			w.Init(os.Stdout, 0, 4, 4, ' ', 0)
			fmt.Fprintln(w, "Id\tDomain name\tRegistrant\tExpiry")
			for _, domain := range response.Return.Domains {
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", domain.ID, domain.Domain, domain.RegName, domain.DateBilledUntil)
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
}
