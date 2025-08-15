package dns

import (
	"context"

	"github.com/sitehostnz/gosh/pkg/api"
	goshDns "github.com/sitehostnz/gosh/pkg/api/dns"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deleteCmd represents the domainAdd command
var deleteCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove a domain name",
	RunE: func(cmd *cobra.Command, args []string) error {
		domainName := cmd.Flag("domain").Value.String()
		client := goshDns.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))
		ctx := context.Background()
		domainGetResponse, err := client.GetZone(ctx, goshDns.GetZoneRequest{DomainName: domainName})
		if err != nil {
			return err
		}

		if !domainGetResponse.Status {
			return nil
		}

		_, err = client.DeleteZone(ctx, goshDns.DeleteZoneRequest{DomainName: domainName})
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {

	deleteCmd.Flags().StringP("domain", "d", "", "The domain name to use")
	deleteCmd.MarkFlagRequired("domain")
}
