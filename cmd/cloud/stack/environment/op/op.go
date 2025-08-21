package onepassword

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Cmd represents the domain command
var Cmd = &cobra.Command{
	Use:   "op",
	Short: "Commands moving environment variables from 1Password to SiteHost and vice versa",
}

func init() {

	Cmd.PersistentFlags().StringP("vault", "", "", "The 1Password vault")
	Cmd.MarkPersistentFlagRequired("vault")

	Cmd.PersistentFlags().String("item", "", "The name of the item in 1Password")
	Cmd.MarkPersistentFlagRequired("item")

	Cmd.PersistentFlags().StringP("service-account", "", "", "1password service account")
	Cmd.MarkPersistentFlagRequired("service-account")

	viper.BindPFlag("service-account", Cmd.PersistentFlags().Lookup("service-account"))

}
