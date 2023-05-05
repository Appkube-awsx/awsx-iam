package iamcmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-iam/authenticator"
	"github.com/Appkube-awsx/awsx-iam/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/spf13/cobra"
)

// getConfigDataCmd represents the getConfigData command
var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		vaultUrl := cmd.Parent().PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.Parent().PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.Parent().PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.Parent().PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.Parent().PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.Parent().PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.Parent().PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticator.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)
		// print(authFlag)
		// authFlag := true
		if authFlag {
			userName, _ := cmd.Flags().GetString("userName")
			if userName != "" {
				getClusterDetails(region, crossAccountRoleArn, acKey, secKey, userName, externalId)
			} else {
				log.Fatalln("userName not provided. Program exit")
			}
		}
	},
}

func getClusterDetails(region string, crossAccountRoleArn string, accessKey string, secretKey string, userName string, externalId string) *iam.GetUserOutput {
	log.Println("Getting aws user data")
	listClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)
	input := &iam.GetUserInput{
		UserName: aws.String(userName),
	}
	userDetailsResponse, err := listClient.GetUser(input)
	log.Println(userDetailsResponse.String())
	if err != nil {
		log.Fatalln("Error:", err)
	}
	return userDetailsResponse
}

func init() {
	GetConfigDataCmd.Flags().StringP("userName", "t", "", "user name")

	if err := GetConfigDataCmd.MarkFlagRequired("userName"); err != nil {
		fmt.Println(err)
	}
}
