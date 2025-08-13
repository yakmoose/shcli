package dns

import (
	"github.com/spf13/cobra"
)

// DnsCmd represents the domain command
var RecordCmd = &cobra.Command{
	Use:   "record",
	Short: "Commands for managing dns zone records",
}

func init() {
	RecordCmd.AddCommand(dnsRecordsListCmd)
	RecordCmd.AddCommand(dnsRecordAddCmd)

	RecordCmd.PersistentFlags().StringP("domain", "d", "", "The domain zone to use")
	RecordCmd.MarkFlagRequired("domain")
}
