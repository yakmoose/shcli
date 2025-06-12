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
	"log"
	"reflect"
	"sort"

	"strings"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cloudDbUserAttach represents the domainAdd command
var cloudDbUserDetach = &cobra.Command{
	Use:   "detach",
	Short: "Detach a user from all  databases on a given server, only if the grants are the same",
	RunE: func(cmd *cobra.Command, args []string) error {

		ctx := context.Background()
		api := api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId"))

		grantClient := grant.New(api)
		databaseClient := db.New(api)
		userClient := user.New(api)

		grants := strings.Split(cmd.Flag("grants").Value.String(), ",")
		sort.Strings(grants)

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

		log.Printf("[DEBUG] Detaching user: %s with grants: %s", userResponse.User.Username, grants)

		if nil != err {
			return err
		}

		for _, database := range databaseListResponse.Return.Databases {

			log.Printf("[DEBUG] Detaching user: %s from database: %s", userResponse.User.Username, database.DBName)

			if !helper.Has(userResponse.User.Grants, func(g models.Grant) bool {
				return g.DBName == database.DBName
			}) {
				log.Printf("[DEBUG] Skipping deatching user: %s to database: %s (grant does not exist)", userResponse.User.Username, database.DBName)
				continue
			}

			g := helper.First(userResponse.User.Grants, func(g models.Grant) bool {
				return g.DBName == database.DBName
			})
			sort.Strings(g.Grants)

			if !reflect.DeepEqual(g.Grants, grants) {
				log.Printf("[DEBUG] Skipping deatching user: %s to database: %s (%s %s)", userResponse.User.Username, database.DBName, g.Grants, grants)
				continue
			}

			response, err := grantClient.Delete(ctx, grant.DeleteRequest{
				MySQLHost:  cmd.Flag("host").Value.String(),
				ServerName: cmd.Flag("server").Value.String(),
				Username:   cmd.Flag("user").Value.String(),
				Database:   database.DBName,
			})

			if nil != err {
				return err
			}

			log.Printf("[DEBUG] Waiting for detaching user: %s from database: %s (grants match)", userResponse.User.Username, database.DBName)

			// ideally we need/want to do these all at once, but locking and stuff...
			return helper.WaitForAction(api, job.GetRequest{ID: response.Return.Job.ID, Type: response.Return.Job.Type})

		}
		return nil
	},
}

func init() {
	cloudDbUserCommand.AddCommand(cloudDbUserDetach)

	cloudDbUserDetach.Flags().StringP("server", "S", "", "The server name")
	cloudDbUserDetach.MarkFlagRequired("server")

	cloudDbUserDetach.Flags().StringP("host", "H", "", "The database host")
	cloudDbUserDetach.MarkFlagRequired("host")

	cloudDbUserDetach.Flags().StringP("user", "U", "", "The database user")
	cloudDbUserDetach.MarkFlagRequired("user")

	cloudDbUserDetach.Flags().StringP("grants", "g", "", "The database user")
	cloudDbUserDetach.MarkFlagRequired("grants")
}
