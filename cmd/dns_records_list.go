package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sitehostnz/gosh/pkg/models"
	"github.com/sitehostnz/gosh/pkg/net"
	"github.com/sitehostnz/terraform-provider-sitehost/sitehost/helper"
	"os"
	"text/tabwriter"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/dns"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listRecordsCmd = &cobra.Command{
	Use:   "list",
	Short: "List zones records",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := dns.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))
		domainName := cmd.Flag("domain").Value.String()

		records, err := client.ListRecords(context.Background(), dns.ListRecordsRequest{Domain: domainName})
		if err != nil {
			return err
		}

		format := cmd.Flag("format").Value.String()

		if cmd.Flag("rr_type").Value.String() != "" {
			rrType := cmd.Flag("rr_type").Value.String()
			records.Return = helper.Filter(records.Return, func(record models.DNSRecord) bool {
				return record.Type == rrType
			})
		}

		if format == "json" {
			json, err := json.MarshalIndent(records, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(json))
		} else if format == "text" {
			w := new(tabwriter.Writer)
			w.Init(os.Stdout, 0, 4, 4, ' ', 0)
			fmt.Fprintln(w, "Id\tDomain\tName\tType\tPriority\tValue")
			for _, record := range records.Return {
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n", record.ID, record.Domain, net.DeconstructFqdn(record.Name, record.Domain), record.Type, record.Priority, record.Content)
			}

			fmt.Fprintln(w)
		} else {
			return errors.New("unsupported format, please choose text or json")
		}

		return nil
	},
}

func init() {
	recordCmd.AddCommand(listRecordsCmd)
	listRecordsCmd.Flags().StringP("rr_type", "t", "", "Filter record RR types")

}
