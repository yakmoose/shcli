/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/bandwidth"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"text/tabwriter"
)

// bandwidthIpAddressCmd fetches the list of IP addresses for the account
var bandwidthIpAddressCmd = &cobra.Command{
	Use:   "ip",
	Short: "List IP addresses",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := bandwidth.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))
		ipAddressesResponse, err := client.ListIPAddresses(context.Background())
		if err != nil {
			return err
		}

		format := cmd.Flag("format").Value.String()
		if format == "json" {
			json, err := json.MarshalIndent(ipAddressesResponse.Return, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(json))
		} else if format == "text" {
			w := new(tabwriter.Writer)
			w.Init(os.Stdout, 0, 4, 4, ' ', 0)
			fmt.Fprintln(w, "Address\tFamily\tNetmask\tPrefix\tRDNS")
			for _, address := range ipAddressesResponse.Return {
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", address.IP, address.Family, address.Netmask, address.Prefix, address.RDNS)
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
	bandwidthCommand.AddCommand(bandwidthIpAddressCmd)
}
