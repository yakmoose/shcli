/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/dns"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// recordAddCmd represents the domainAdd command
var recordAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new domain record",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := dns.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))
		ctx := context.Background()

		domainName := cmd.Flag("domain").Value.String()

		response, err := client.AddRecord(ctx, dns.AddRecordRequest{
			Domain:   domainName,
			Type:     cmd.Flag("rr_type").Value.String(),
			Name:     cmd.Flag("name").Value.String(),
			Content:  cmd.Flag("content").Value.String(),
			Priority: cmd.Flag("priority").Value.String(),
		})

		if err != nil {
			return err
		}

		json, err := json.MarshalIndent(response, "", "  ")
		fmt.Println(string(json))

		return nil
	},
}

func init() {
	recordCmd.AddCommand(recordAddCmd)
	recordAddCmd.Flags().StringP("domain", "d", "", "The domain name to use")
	recordAddCmd.MarkFlagRequired("domain")

	recordAddCmd.Flags().StringP("rr_type", "t", "", "The rr type name to use")
	recordAddCmd.MarkFlagRequired("rr_type")

	recordAddCmd.Flags().StringP("name", "n", "", "The rr type name to use")
	recordAddCmd.MarkFlagRequired("name")

	recordAddCmd.Flags().StringP("content", "c", "", "The record content")
	recordAddCmd.MarkFlagRequired("content")

	recordAddCmd.Flags().StringP("priority", "p", "", "The record priority")
}
