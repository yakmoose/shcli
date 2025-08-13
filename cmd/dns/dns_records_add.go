package dns

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sitehostnz/gosh/pkg/api"
	goshDns "github.com/sitehostnz/gosh/pkg/api/dns"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// dnsRecordAddCmd represents the domainAdd command
var dnsRecordAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new domain record",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := goshDns.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))
		ctx := context.Background()

		domainName := cmd.Flag("domain").Value.String()

		response, err := client.AddRecord(ctx, goshDns.AddRecordRequest{
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

	dnsRecordAddCmd.Flags().StringP("domain", "d", "", "The domain name to use")
	dnsRecordAddCmd.MarkFlagRequired("domain")

	dnsRecordAddCmd.Flags().StringP("rr_type", "t", "", "The rr type name to use")
	dnsRecordAddCmd.MarkFlagRequired("rr_type")

	dnsRecordAddCmd.Flags().StringP("name", "n", "", "The rr type name to use")
	dnsRecordAddCmd.MarkFlagRequired("name")

	dnsRecordAddCmd.Flags().StringP("content", "c", "", "The record content")
	dnsRecordAddCmd.MarkFlagRequired("content")

	dnsRecordAddCmd.Flags().StringP("priority", "p", "", "The record priority")
}
