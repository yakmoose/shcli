/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// stackGetDockerFileCmd represents the get command
var stackGenerateDockerFileCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate the dockerfile",
	RunE: func(cmd *cobra.Command, args []string) error {

		//client := image.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))
		//
		////serverName := cmd.Flag("server").Value.String()
		////stackName := cmd.Flag("stack").Value.String()
		//imageName := cmd.Flag("image").Value.String()
		//
		//stackResponse, err := client.List(context.Background())
		//if err != nil {
		//	return err
		//}
		//
		//// fmt.Println(stackResponse.Stack.DockerFile)

		return nil
	},
}

func init() {
	stackImageCmd.AddCommand(stackGenerateDockerFileCmd)

	//stackGetDockerFileCmd.Flags().StringP("server", "S", "", "The server name to fetch the stack list from")
	//stackGetDockerFileCmd.MarkFlagRequired("server")
	//
	//stackGetDockerFileCmd.Flags().StringP("stack", "s", "", "The stack to get")
	//stackGetDockerFileCmd.MarkFlagRequired("stack")

	stackGenerateDockerFileCmd.Flags().String("image", "", "The name of the image")
	stackGetDockerFileCmd.MarkFlagRequired("image")
}
