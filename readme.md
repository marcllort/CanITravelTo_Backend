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
To make the backend always available, we need to create a system.d service. In the path `cd /etc/systemd/system/`
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

When running the *Gin router* previously I used `run`, which serves HTTP. But since being deployed I use `runTLS`, which servers HTTPS. In this case you need to provide two certificates, later explained in Domains and Cloudfare.

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

## Domains and Cloudfare

Currently two domains are being used, canitravelto.com (where the frontend is hosted, Github Pages) and canitravelto.wtf (where the backend API is hosted, EC2 AWS).

Both are using Cloudfare, which provides caching for the website, important for the frontend as Github offers a maximum of 100GB of bandwidth per month, but most importantly it provides TLS certificates, so the website and backend are HTTPS encrypted and safe to use.

At first, only the frontend used Cloudfare as there was no need for the backend to use HTTPS, until I saw that an HTTPS website can't consume from a HTTP API. The options where to go back to HTTP in the frontend (I didn't manage to do it, because Github Pages always provides a HTTPS).

The second option was to serve the API as HTTPS. I created my own certificates, but as they were self-signed, HTTPS didn't like them, so I had to have valid certificates. To get a SSL certificate, you need a domain name, so I then acquired canitravelto.wtf, to use it instead of the public AWS IP.

[canitravelto.com](canitravelto.com) is hosted by godaddy.com

[canitravelto.wtf](canitravelto.wtf) is hosted by name.com

### Cloudfare
To configure Cloudfare in both domains, there's a few steps to follow:

#### Frontend
In the case of Github Pages (canitravelto.com) just follow the Cloudfare set-up. Once the email is received about your website being active, navigate to SSL/TLS and change the mode to FULL. If not done, the webpage won't be reached (still don't know why)

Another change to be done, is CORS in the API requests. CORS stands for Cross-origin resource sharing, which means that it will consume resources from another site. If not correctly configured, the requests will fail.

To make this work, in the frontend we only need to add to the header of the request this two lines:

 `myHeaders.append("Access-Control-Allow-Origin", "*");`
    
 `myHeaders.append("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers");`


#### Backend
With the backend (canitravelto.wtf), first I hosted in [AWS with Route53](https://www.youtube.com/watch?v=qor31Egu0Rg) (probably not needed) and then did the same configuration as with Github pages in Cloudfare.

In this case we also have to navigate to SSL/TLS -> Origin Server and select Create Server. We need to copy the two keys and save them in "http-server.key" and "http-server.cert". 

In the GoLang backend (*main.go*):

`router.RunTLS(PORT, "Creds/https-server.crt", "Creds/https-server.key")`

We also need to enable CORS on the backend, to enable the requests from the frontend. Just add this two headers to the returned JSON:

`c.Header("Access-Control-Allow-Origin", "*")`
    
`c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")`


go mod init
go mod tidy
Performance comparison, in local and in web

## Expiring Dates
   
   All logins with personal email
   
   - canitravelto.com godaddy.com (May 2021) 
   - AWS RDS MySQL BBDD (June 2021) 
   - AWS EC2 Ubuntu t2.micro Instance (June 2021) 
   - canitravelto.wtf - name.com (June 2021) 
   - Cloudfare check just in case (June 2021) 
   - Cloudfare SSL for .com and .wtf (June 2035)

## TO-DO
  
  - [ ] Add Covid info to response, and modify responses to be already displayed in frontend
  - [ ] Card view for backend response to frontend (show covid cases, visa status...)
  - [ ] Create Postman test scenarios  
  - [ ] Commentate code 
  - [ ] Retrieve accesses to the website (IP, origin, destination...) or find solution to log visits/web usage (gin may already have one)
  - [ ] Suggestions in Frontend (i.e You can also go there with your passport!) 
  - [ ] Unify country names between frontend, passportInfo and CovidInfo (maybe new row in Covid info with name of PassportInfo?)
  - [ ] Add travis-ci or gitlab pipeline and update link of travis-ci build in readme.md
  - [ ] Improve deployment to github pages and ubuntu improve script
  - [ ] Remove countries.go list and just sanitise input to prevent sqlInjection
  - [ ] Add tests to golang and react
  - [ ] Develop good frontend
  - [ ] Protect ip for backend (cors, autotls?)
  - [ ] Make Frontend use GET endpoint
  - [ ] Archive old frontend repository
  - [ ] Save user emails from frontend form
  - [ ] Write frontend README and update REACT part of this one (now use html,css)
  - [ ] Google SEO
  - [ ] Move out from personal mail
  - [ ] Alternative to AWS? 11 months remaining (June 2021)
  - [x] Add cloudfare
  - [x] Change Domain from Amazon to github
  

