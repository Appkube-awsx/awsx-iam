package cmd

import (
	"log"
	"os"

	"github.com/Appkube-awsx/awsx-iam/authenticator"
	"github.com/Appkube-awsx/awsx-iam/client"
	"github.com/Appkube-awsx/awsx-iam/cmd/iamcmd"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/spf13/cobra"
)

var AwsxUserMetadataCmd = &cobra.Command{
	Use:   "getListClusterMetaDataDetails",
	Short: "getListClusterMetaDataDetails command gets resource counts",
	Long:  `getListClusterMetaDataDetails command gets resource counts details of an AWS account`,

	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Command iam for user started")
		vaultUrl := cmd.PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticator.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)

		if authFlag {
			listUserData(region, crossAccountRoleArn, acKey, secKey, externalId)
		}
	},
}



// json.Unmarshal
func listUserData(region string, crossAccountRoleArn string, accessKey string, secretKey string, externalId string) (*iam.ListUsersOutput , error) {
	log.Println("getting user metadata list summary")

	listClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)
	Request := &iam.ListUsersInput{}
	Response, err := listClient.ListUsers(Request)
	if err != nil {
		log.Fatalln("Error:in getting  user list", err)
	}
	log.Println(Response)
	return Response, err 
	
}


func Execute() {
	err := AwsxUserMetadataCmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		os.Exit(1)
	}
}

func init() {
	
	AwsxUserMetadataCmd.AddCommand(iamcmd.GetConfigDataCmd)
	
	AwsxUserMetadataCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxUserMetadataCmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxUserMetadataCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxUserMetadataCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxUserMetadataCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxUserMetadataCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws crossAccountRoleArn is required")
	AwsxUserMetadataCmd.PersistentFlags().String("externalId", "", "aws external id auth")

}
