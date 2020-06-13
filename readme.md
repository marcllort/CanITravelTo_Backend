# CanITravelTo

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

## Database

[Amazon RDS MySQL](https://aws.amazon.com/es/rds/mysql/) free instance (t.2 micro instance) --> Basic/Simple setup (hosted in Paris). 

Set its configuration to publicly available.

In Security groups, add 2 inbound rules (port 3306 mysql), one for your development computer with your own IP, and another for the EC2 instances where the backend server is hosted. No need to set the IP, just the name of its security group/launch-wizard number 

The current dataset in the DB has the information from [PassportIndex](passportindex.com) of the places you can travel with your passport. 

This dataset, which should be regularly updated, can be found in [GitHub](https://github.com/ilyankou/passport-index-dataset) in CSV.Transform it to MySQL [here](https://www.convertcsv.com/csv-to-sql.htm).

If updated, and the list of countries has changed, it must also be changed in Countries.go list

Once the import script is prepared, just connect to the DB with DataGrip/Workbench and run the script.

The DB credentials should be stored always in the Creds folder in the Backend project, with the following format:
```json
    {
      "user": "admin",
      "password": "password set in AWS",
      "hostname": "x.x.eu-west-3.rds.amazonaws.com (endpoint field in AWS)",
      "port": "3306",
      "database": "db name"
    }
```

## EC2 Ubuntu

First I'll explain how to create the virtual-machine where the backend will run, and later in the *Backend GoLang* I will explain the backend itself.

Hosted in free-tier t2.micro [Amazon EC2 instance](https://aws.amazon.com/es/ec2/), running Ubuntu 18.04 default configuration (hosted in Paris).

When creating the instance, download the keypair.pem to be able to SSH into the machine.

The only configuration needed, is in Security groups, where there's the need to, in the inbound rules, the ports 22 (SSH), 80 (HTTP), 8080 (DEV) and 443 (HTTPS) should be open to "Anywhere", so to `0.0.0.0`.

##### To ssh into AWS Ubuntu machine: 
`ssh -i keypair.pem ubuntu@[Public-DNS] (i.e= whatever.eu-west-3.compute.amazonaws.com)`

##### Compile to AWS:
If not already done (to check do: `go env`), do a:
```sh
set GOOS=linux
set GOARCH=amd64 
```
And to compile the new version:

```sh
go build -o canITravelToUpdated
```


##### To upload the new build:
`scp -i keypair.pem fileToUpload ubuntu@[Public-DNS]: (i.e= whatever.eu-west-3.compute.amazonaws.com:)`

##### Create service in Ubuntu
To make the backend always available, we need to create a system.d service.
The file, should follow this pattern:

    [Unit]
    Description=CanITravelTo Backend
    
    [Service]
    ExecStart=/home/<username>/<exepath> [params]
    User=root
    Group=root
    Restart=always
    
    [Install]
    WantedBy=multi-user.target

To start the service:
```sh
sudo systemctl enable <filename>.service
sudo systemctl start <filename>.service
sudo systemctl status <filename>.service
```
To stop the service:
```sh
sudo systemctl stop <filename>.service
```

To see the backend logs in a "tail" way:
```sh
sudo journalctl --follow _SYSTEMD_UNIT=canitravelto.service
```
#### Custom Domain name
To configure the ROUTE54 I used: https://www.youtube.com/watch?v=qor31Egu0Rg

## Backend GoLang

Backend is written in Go, and using *Gin framework* for the http requests. The connection with the mySQL DB is done with *go-sql-driver*.

When working on local, you should connect to *localhost:8080/travel*. If testing with the hosted backend in EC2, *publicIP:8080/travel*.

##### Libraries used:
  - Gin
  - go-sql-driver

The request to the backend should always be a *POST*, and this could be an example JSON body for the request:
```json
    {
        "destination": "Spain",
        "origin": "France"
    }
```

So far, the response time in local is about 1ms, while in AWS is around 36ms. In both cases has been stress tested with thousands of requests every 1ms, and has been able to not drop a single request.

To stress test the backend, I used the chrome extension named [RestfulStress](https://chrome.google.com/webstore/detail/restful-stress/lljgneahfmgjmpglpbhmkangancgdgeb).

## Frontend React

The frontend is a static website coded in React and hosted in Github Pages, which is free with a maximum of a 100GB of bandwidth per month. To avoid this limitation, Cloudfare can be used. Cloudfare will (for free) cache the website in their servers and also provide a Secure SSL certificate. To do so, follow [this](https://www.toptal.com/github/unlimited-scale-web-hosting-github-pages-cloudflare) tutorial.

To create and build the project, [*npm*](https://www.npmjs.com/) and [*node.js*](https://nodejs.org/en/) will be needed.

Install *create-react-app*:

`npm i create-react-app`

And create the project:

`npx create-react-app my-app`

We need to install the Github-pages dependency, so we can deploy the web:

`npm install gh-pages --save-dev`

Go to the *package.json* file, and add: (or whatever domain name you want to use)
```json
"homepage": "https://{username}.github.io/{repo-name}"
```

Also add in the begining of the scripts section:
```json
"scripts": {
    ...
    "predeploy": "npm run build",
    "deploy": "gh-pages -d build"
}
```

Then just code! To deploy to Github pages:

`npm run deploy`

The only thing left to do is: go to the repository->Settings->Github Pages->Source and select: *gh-pages* branch.

Then, set the domain to: *canitravelto.com* (This step of setting the domain for now must be done every time there is a new deployment, because the CNAME file is deleted with the deployment)

For this deployment I used [this](https://dev.to/yuribenjamin/how-to-deploy-react-app-in-github-pages-2a1f) and [this](https://medium.com/@shauxna/setting-up-a-custom-domain-for-your-react-app-on-github-pages-827b2606ca18) tutorial.

## TO-DO
  - [ ] Remove countries.go list and just sanitise input to prevent sqlInjection
  - [ ] Add tests
  - [ ] Add travis-ci pipeline and update link of travis-ci build in readme.md
  - [ ] Move out from personal mail
  - [ ] Add cloudfare
  - [ ] Protect ip for backend
  - [ ] Alternative to AWS? 11 months remaining
  - [x] Change Domain from Amazon to github
