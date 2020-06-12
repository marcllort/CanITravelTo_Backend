**CanITravelTo**

Libraries used:
Gin
mysql-connector


Passport Info retrieval: (passportindex.com)
`https://github.com/ilyankou/passport-index-dataset`

Converted to mysql with: 
`https://www.convertcsv.com/csv-to-sql.htm`

If updated, and the list of countries has changed, it must also be changed in Countries.go list (or remove that list and just sanitise input to prevent sqlInjection)

Compile to AWS:
If not already done (to check do: `go env`), do a set GOOS=linux and set GOARCH=amd64 
`go build -o canITravelToUpdated`

To upload the new build:
`scp -i Creds/key-aws.pem CanITravelToUpdated ubuntu@ec2-35-180-85-2.eu-west-3.compute.amazonaws.com:`

To ssh into AWS Ubuntu machine:
`ssh -i Creds/key-aws.pem ubuntu@ec2-35-180-85-2.eu-west-3.compute.amazonaws.com`


To see the logs in a "tail" way:

`sudo journalctl --follow _SYSTEMD_UNIT=canitravelto.service`

