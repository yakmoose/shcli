package contacts

import (
	"fmt"
	"os"
	"slices"
	"text/tabwriter"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/srs"
	"github.com/sitehostnz/gosh/pkg/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"context"
	"encoding/json"
	"errors"
)

var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all domain name contacts",
	RunE: func(cmd *cobra.Command, args []string) error {

		client := srs.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))

		response, err := client.ListContacts(context.Background())
		if err != nil {
			return err
		}

		format := cmd.Flag("format").Value.String()
		if format == "json" {
			json, err := json.MarshalIndent(response.DomainContacts, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(json))
		} else if format == "text" {
			w := new(tabwriter.Writer)
			w.Init(os.Stdout, 0, 4, 4, ' ', 0)
			fmt.Fprintln(w, "Id\tName\tEmail\tNumber of domains")
			slices.SortFunc(response.DomainContacts, func(a models.DomainContact, b models.DomainContact) int {
				return int(b.DomainCount) - int(a.DomainCount)
			})
			for _, contact := range response.DomainContacts {
				fmt.Fprintf(w, "%d\t%s\t%s\t%d\n", contact.ID, contact.Name, contact.Email, contact.DomainCount)
			}

			fmt.Fprintln(w)
		} else {
			/// how to error out.
			return errors.New("unsupported format, please choose text or json")
		}

		return nil
	},
}
