/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"context"
	"github.com/sitehostnz/gosh/pkg/api/cloud/db"
	"github.com/sitehostnz/gosh/pkg/api/cloud/db/grant"
	"github.com/sitehostnz/gosh/pkg/api/cloud/db/user"
	"github.com/sitehostnz/gosh/pkg/api/job"
	"github.com/sitehostnz/gosh/pkg/models"
	"github.com/sitehostnz/terraform-provider-sitehost/sitehost/helper"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"strings"
)

// cloudDbUserAttach represents the domainAdd command
var cloudDbUserAttach = &cobra.Command{
	Use:   "attach",
	Short: "Attach a user to all databases on a given server with the supplied grants",
	RunE: func(cmd *cobra.Command, args []string) error {

		ctx := context.Background()
		api := api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId"))

		grantClient := grant.New(api)
		databaseClient := db.New(api)
		userClient := user.New(api)

		grants := strings.Split(cmd.Flag("grants").Value.String(), ",")

		databaseListResponse, err := databaseClient.List(ctx, db.ListOptions{
			MySQLHost:  cmd.Flag("host").Value.String(),
			ServerName: cmd.Flag("server").Value.String(),
		})

		if nil != err {
			return err
		}

		userResponse, err := userClient.Get(ctx, user.GetRequest{
			MySQLHost:  cmd.Flag("host").Value.String(),
			ServerName: cmd.Flag("server").Value.String(),
			Username:   cmd.Flag("user").Value.String(),
		})

		if nil != err {
			return err
		}

		for _, database := range databaseListResponse.Return.Databases {

			log.Printf("[DEBUG] Attaching user: %s to database: %s", userResponse.User.Username, database.DBName)

			// if the user is already attached grants, we will skip...
			if helper.Has(userResponse.User.Grants, func(g models.Grant) bool {
				return g.DBName == database.DBName
			}) {
				log.Printf("[DEBUG] Skipping attaching user: %s to database: %s (grant exists)", userResponse.User.Username, database.DBName)
				continue
			}

			response, err := grantClient.Add(ctx, grant.AddRequest{
				MySQLHost:  cmd.Flag("host").Value.String(),
				ServerName: cmd.Flag("server").Value.String(),
				Username:   cmd.Flag("user").Value.String(),
				Grants:     grants,
				Database:   database.DBName,
			})

			if nil != err {
				return err
			}

			// ideally we need/want to do these all at once, but locking and stuff...
			log.Printf("[DEBUG] Waitinf for attaching user: %s to database: %s", userResponse.User.Username, database.DBName)

			return helper.WaitForAction(api, job.GetRequest{ID: response.Return.Job.ID, Type: response.Return.Job.Type})
		}

		return nil
	},
}

func init() {
	cloudDbUserCommand.AddCommand(cloudDbUserAttach)

	cloudDbUserAttach.Flags().StringP("server", "S", "", "The server name")
	cloudDbUserAttach.MarkFlagRequired("server")

	cloudDbUserAttach.Flags().StringP("host", "H", "", "The database host")
	cloudDbUserAttach.MarkFlagRequired("host")

	cloudDbUserAttach.Flags().StringP("user", "U", "", "The database user")
	cloudDbUserAttach.MarkFlagRequired("user")

	cloudDbUserAttach.Flags().StringP("grants", "g", "", "The database user")
	cloudDbUserAttach.MarkFlagRequired("grants")

}
