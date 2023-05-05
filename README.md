# IAM CLi's

## To list all the IAM function,run the following command:

```bash
awsx-iam --zone <zone> --acccessKey <acccessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn> --externalId <externalId>
```

## To retrieve the configuration details of a specific IAM function in iamcmd, run the following command:

```bash
awsx-iam getConfigData -t <table> --zone <zone> --accessKey <accessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn> --external <externalId>  --userName <userName>
```
