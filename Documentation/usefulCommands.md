To ssh into AWS Ubuntu machine:

`ssh -i Creds/key-aws.pem ubuntu@ec2-35-180-85-2.eu-west-3.compute.amazonaws.com`

Then, just run update.sh, and the latest git changes will be applied

To check the logs (inside the CanITravelTo_Backend folder):
`docker-compose logs -f`

------------------------------------------------------------------------------------------------------------------------

Compile to AWS [DEPRECATED]:

If not already done (to check do: `go env`), do a set GOOS=linux and set GOARCH=amd64 
`go build -o canITravelToUpdated`

To upload the new build:
`scp -i Creds/key-aws.pem CanITravelToUpdated ubuntu@ec2-35-180-85-2.eu-west-3.compute.amazonaws.com:`

To see the logs in a "tail" way:

`sudo journalctl --follow _SYSTEMD_UNIT=canitravelto.service`

Endpoint for backend is: 
http://35.180.85.2:8080/travel

If not already done (to check do: `go env`), do a:
```sh
set GOOS=linux
set GOARCH=amd64 
```
And to compile the new version:

```sh
go build -o canITravelToUpdated
```
