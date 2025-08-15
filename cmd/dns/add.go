package dns

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/sitehostnz/gosh/pkg/api"
	goshDns "github.com/sitehostnz/gosh/pkg/api/dns"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the domainAdd command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new domain name",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := goshDns.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))
		ctx := context.Background()

		domainName := cmd.Flag("domain").Value.String()

		domainResponse, err := client.GetZone(ctx, goshDns.GetZoneRequest{DomainName: domainName})
		if err != nil {
			return err
		}

		if domainResponse.Status {
			return errors.New("Domain already exists")
		}

		zoneCreateResponse, err := client.CreateZone(ctx, goshDns.CreateZoneRequest{DomainName: domainName})
		if err != nil {
			return err
		}

		json, err := json.MarshalIndent(zoneCreateResponse.Return, "", "  ")
		fmt.Println(string(json))

		return nil
	},
}

func init() {
	addCmd.Flags().StringP("domain", "d", "", "The domain name to use")
	addCmd.MarkFlagRequired("domain")

}
