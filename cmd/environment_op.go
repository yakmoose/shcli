/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// environmentOpCmd represents the domain command
var environmentOpCmd = &cobra.Command{
	Use:   "op",
	Short: "Commands moving environment variables from 1Password to SiteHost and vice versa",
}

func init() {
	environmentCmd.AddCommand(environmentOpCmd)
	environmentOpCmd.PersistentFlags().StringP("vault", "", "", "The 1Password vault")
	environmentOpCmd.MarkPersistentFlagRequired("vault")

	environmentOpCmd.PersistentFlags().String("item", "", "The name of the item in 1Password")
	environmentOpCmd.MarkPersistentFlagRequired("item")

	environmentOpCmd.PersistentFlags().StringP("service-account", "", "", "1password service account")
	environmentOpCmd.MarkPersistentFlagRequired("service-account")

	viper.BindPFlag("service-account", environmentOpCmd.PersistentFlags().Lookup("service-account"))

}
