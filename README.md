- [What is awsx-iam](#awsx-iam)
- [How to write plugin subcommand](#how-to-write-plugin-subcommand)
- [How to build / Test](#how-to-build--test)
- [what it does ](#what-it-does)
- [command input](#command-input)
- [command output](#command-output)
- [How to run ](#how-to-run)

# awsx-iam

This is a plugin subcommand for awsx cli ( https://github.com/Appkube-awsx/awsx#awsx ) cli.

For details about awsx commands and how its used in Appkube platform , please refer to the diagram below:

![alt text](https://raw.githubusercontent.com/AppkubeCloud/appkube-architectures/main/LayeredArchitecture-phase2.svg)

This plugin subcommand will implement the Apis' related to IAM services , primarily the following API's:

- getConfigData

This cli collect data from metric / logs / traces of the IAM services and produce the data in a form that Appkube Platform expects.

This CLI , interacts with other Appkube services like Appkube vault , Appkube cloud CMDB so that it can talk with cloud services as
well as filter and sort the information in terms of product/services, so that Appkube platform gets the data that it expects from the cli.

# How to write plugin subcommand

Please refer to the instaruction -
https://github.com/Appkube-awsx/awsx#how-to-write-a-plugin-subcommand

It has detailed instruction on how to write a subcommand plugin , build/test/debug/publish and integrate into the main commmand.

# How to build / Test

            go run main.go
                - Program will print Calling aws-cloudelements on console

            Another way of testing is by running go install command
            go install
            - go install command creates an exe with the name of the module (e.g. awsx-iam) and save it in the GOPATH
            - Now we can execute this command on command prompt as below
           awsx-iam getConfigData --zone=us-east-1 --accessKey=xxxxxxxxxx --secretKey=xxxxxxxxxx --crossAccountRoleArn=xxxxxxxxxx  --externalId=xxxxxxxxxx

# what it does

This subcommand implement the following functionalities -
getConfigData - It will get the resource count summary for a given AWS account id and region.

# command input

1. --valutURL = URL location of vault - that stores credentials to call API
2. --acountId = The AWS account id.
3. --zone = AWS region
4. --accessKey = Access key for the AWS account
5. --secretKey = Secret Key for the Aws Account
6. --crossAccountRoleArn = Cross Acount Rols Arn for the account.
7. --external Id = The AWS External id.
8. --userName= Insert your user name which is in aws account.

# command output

Users: [
{
Arn: "arn:aws:iam::657907747545:user/papu",
CreateDate: 2018-02-13 06:27:49 +0000 UTC,
PasswordLastUsed: 2023-05-03 08:41:07 +0000 UTC,
Path: "/",
UserId: "AIDAI24V3LIJXECYEYLHM",
UserName: "papu"
},

{
Arn: "arn:aws:iam::657907747545:user/a.khan",
CreateDate: 2022-12-26 07:13:52 +0000 UTC,
PasswordLastUsed: 2023-05-10 04:27:37 +0000 UTC,
Path: "/",
UserId: "AIDAZSLS3RLMQI7H4SEVP",
UserName: "a.khan"
},
{
Arn: "arn:aws:iam::657907747545:user/a.mustafa",
CreateDate: 2022-12-28 08:02:14 +0000 UTC,
PasswordLastUsed: 2023-05-09 08:03:34 +0000 UTC,
Path: "/",
UserId: "AIDAZSLS3RLM2MGGNRWDE",
UserName: "a.mustafa"
},
]

# How to run

From main awsx command , it is called as follows:

```bash
awsx-iam  --zone=us-east-1 --accessKey=<> --secretKey=<> --crossAccountRoleArn=<>  --externalId=<>
```

If you build it locally , you can simply run it as standalone command as:

```bash
go run main.go  --zone=us-east-1 --accessKey=<> --secretKey=<> --crossAccountRoleArn=<> --externalId=<>
```

# awsx-iam

iam extension

# AWSX Commands for AWSX-IAM Cli's :

1. CMD used to get list of iam instance's :

```bash
./awsx-iam --zone=us-east-1 --accessKey=<6f> --secretKey=<> --crossAccountRoleArn=<> --externalId=<>
```

2. CMD used to get Config data (metadata) of AWS iam instances :

```bash
./awsx-iam --zone=us-east-1 --accessKey=<#6f> --secretKey=<> --crossAccountRoleArn=<> --externalId=<> getConfigData --userName=<>
```
