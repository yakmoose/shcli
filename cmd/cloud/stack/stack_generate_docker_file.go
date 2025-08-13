package stack

//import (
//	"github.com/spf13/cobra"
//)
//
//// getDockerFileCmd represents the get command
//var stackGenerateDockerFileCmd = &cobra.Command{
//	Use:   "generate",
//	Short: "generate the dockerfile",
//	RunE: func(cmd *cobra.Command, args []string) error {
//
//		/// imageClient := image.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))
//
//		//serverName := cmd.Flag("server").Value.String()
//		//stackName := cmd.Flag("stack").Value.String()
//		// imageName := cmd.Flag("image").Value.String()
//
//		//imageResponse, err := imageClient.Get(
//		//	context.Background(),
//		//	image.
//		//		Code: imageName,
//		//	},
//		//)
//
//		//json, err := json.MarshalIndent(imageResponse.Image, "", "  ")
//		//if err != nil {
//		//	return err
//		//}
//		//fmt.Println(string(json))
//
//		return nil
//	},
//}
//
//func init() {
//	stackImageCmd.AddCommand(stackGenerateDockerFileCmd)
//
//	//getDockerFileCmd.Flags().StringP("server", "S", "", "The server name to fetch the stack list from")
//	//getDockerFileCmd.MarkFlagRequired("server")
//
//	stackGenerateDockerFileCmd.Flags().String("image", "", "The name of the image")
//	getDockerFileCmd.MarkFlagRequired("image")
//}
