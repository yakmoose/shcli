package stack

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/tabwriter"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/cloud/server"
	"github.com/sitehostnz/gosh/pkg/api/cloud/stack"
	"github.com/sitehostnz/gosh/pkg/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func formatStackOutput(format string, stack models.Stack) error {

	if format == "json" {
		json, err := json.MarshalIndent(stack, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(json))
	} else if format == "text" {

		pattern := "VIRTUAL_HOST=(.*)"

		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 4, 4, ' ', 1)
		fmt.Fprintln(w, "Label\tValue")
		fmt.Fprintf(w, "Label\t%s\n", stack.Label)
		fmt.Fprintf(w, "Name\t%s\n", stack.Name)
		fmt.Fprintf(w, "Server Name\t%s\n", stack.Server)
		fmt.Fprintf(w, "Server Label\t%s\n", stack.ServerLabel)
		expr, _ := regexp.Compile(pattern)
		m := expr.FindStringSubmatch(stack.DockerFile)
		if len(m) == 2 {
			fmt.Fprintf(w, "Aliases\t%s\n", strings.Join(strings.Split(m[1], ","), ", "))
		}
		fmt.Fprintf(w, "Admin URL\thttps://cp.sitehost.nz/cloud/manage-container/server/%s/stack/%s\n", stack.Server, stack.Name)
		fmt.Fprintln(w)

	} else {
		/// how to error out.
		return errors.New("unsupported format, please choose text or json")
	}

	return nil
}

// findCmd represents the list command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find a stack",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId"))
		stackServerClient := server.New(client)
		stackClient := stack.New(client)
		hostPattern := cmd.Flag("host-pattern").Value.String()

		stackServersResponse, err := stackServerClient.List(context.Background())
		if err != nil {
			return err
		}

		pattern := regexp.QuoteMeta(hostPattern)
		expr, _ := regexp.Compile(pattern)

		for _, ss := range stackServersResponse.CloudServers {
			stackResponse, err := stackClient.List(context.Background(), stack.ListRequest{ServerName: ss.Name})
			if err != nil {
				return err
			}

			for _, s := range stackResponse.Return.Stacks {
				// check the name,
				if s.Name == hostPattern || s.Label == hostPattern {
					formatStackOutput(cmd.Flag("format").Value.String(), s)
					continue
				}

				// then check the aliases/vhost stuff
				// get the virtual host and split
				m := expr.Match([]byte(s.DockerFile))
				if m {
					formatStackOutput(cmd.Flag("format").Value.String(), s)
					continue
				}

			}

		}
		return nil
	},
}

func init() {
	findCmd.Flags().StringP("host-pattern", "p", "", "The host name to look for")
	findCmd.MarkFlagRequired("host-pattern")
}
