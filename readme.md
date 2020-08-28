# CanITravelTo

![CI](https://github.com/marcllort/CanITravelTo_Backend/workflows/CI/badge.svg)
![CD](https://github.com/marcllort/CanITravelTo_Backend/workflows/CD/badge.svg)

- [CanITravelTo](#canitravelto)
  * [Motivation](#motivation)
  * [Screenshots](#screenshots)
  * [Database](#database)
  * [EC2 Ubuntu](#ec2-ubuntu)
  * [Backend GoLang](#backend-golang)
    + [Usage](#usage)
    + [CORS](#cors)
    + [Data Retriever](#data-retriever)
    + [Business Handler](#business-handler)
  * [Docker](#docker)
    + [Kubernetes](#kubernetes)
  * [Git](#git)
    + [CI/CD](#cicd)
      - [SSH from pipeline](#ssh-from-pipeline)
  * [Frontend](#frontend)
  * [Domains and Cloudfare](#domains-and-cloudfare)
    + [Cloudfare](#cloudfare)
      - [Frontend](#frontend-1)
      - [Backend](#backend)


## Motivation

The main purpose of this web-service has been learning how to develop a full production-ready environment, using the latest technologies (Go, Docker, AWS, CI/CD, microservices...), 
it's **not supposed to be a "usable" service**, that's why the frontend may look a bit rough.
The service has a simple static website frontend that consumes from the backend through an API (TLS ready, so the website can be https). 

Its purpose is to let the user **know if you can travel to a destination country from your country of origin**. 
The functionality is partially working, as it shows if with the passport of the user, he/she can travel there VISA free, a specific amount of days, VISA required... 

It also shows the amount of Covid cases in the destination country, but the actual travel restrictions are not considered, as to get that information I would have to manually insert it and keep updated, or pay for a premium API. 

Another option would be to scrap the website of each country to find the information automatically, but I'm not that crazy/have that much time (195 countries!).

## Screenshots

[Website (canitravelto.com)](https://canitravelto.com)

![Example Web](https://github.com/marcllort/CanITravelTo_Backend/blob/master/Documentation/Assets/webExample.gif)

[Postman Collection](https://github.com/marcllort/CanITravelTo_Backend/blob/master/Documentation/Postman) call example:

![Example Postman](https://github.com/marcllort/CanITravelTo_Backend/blob/master/Documentation/Assets/postmanExample.gif)

## Database

Running on [Amazon RDS MySQL](https://aws.amazon.com/es/rds/mysql/) free instance (t.2 micro instance) --> Basic/Simple setup (hosted in Paris). 

Steps to configure:

   * Set its configuration to publicly available.

   * In Security groups, add 2 inbound rules (port 3306 mysql), one for your development computer with your own IP, and another for the EC2 instances where the backend server is being hosted. No need to set the IP, just the name of its security group/launch-wizard number.

   * The current dataset in the DB has the information from [PassportIndex](passportindex.com) of the places you can travel with your passport. 

   * This dataset, which should be regularly updated, can be found in [GitHub](https://github.com/ilyankou/passport-index-dataset) in CSV. Transform it to MySQL [here](https://www.convertcsv.com/csv-to-sql.htm). If updated, and the list of countries has changed, it must also be changed in Countries.go list

   * Once the import script is prepared, just connect to the DB with DataGrip/Workbench and run the script.

The DB credentials should be stored always in the Creds folder in the [Backend GoLang](#backend-golang), with the following format:
```json
{
  "user": "admin",
  "hostname": "x.x.eu-west-3.rds.amazonaws.com (endpoint field in AWS)",
  "port": "3306",
  "database": "db name"
}
```

## EC2 Ubuntu 

First I'll explain how to create the virtual-machine where the backend will run, and later in the [Backend GoLang](#backend-golang) I will explain the backend itself.

Hosted in free-tier t2.micro [Amazon EC2 instance](https://aws.amazon.com/es/ec2/), running Ubuntu 18.04 default configuration (hosted in Paris).

When creating the instance, download the keypair.pem to be able to SSH into the machine.

The only configuration needed, is in Security groups, where there's the need to, in the inbound rules, the ports 22 (SSH), 80 (HTTP), 8080 (DEV) and 443 (HTTPS) should be open to "Anywhere", so to `0.0.0.0`.

##### To ssh into AWS Ubuntu machine: 
`ssh -i keypair.pem ubuntu@[Public-DNS] (i.e= whatever.eu-west-3.compute.amazonaws.com)`


## Backend GoLang

Backend is being **written in Go**, and using *Gin framework* for the http requests. The connection with the mySQL DB is done with *go-sql-driver*. Now, there are two microservices, one for data retrieval, and the other one handling the requests.

When working on local, you should connect to `localhost:8080/travel`. If testing with the hosted backend in EC2, `publicIP:8080/travel`.

##### Libraries used:
  - GinGonic
  - go-sql-driver

Using the following commands, go generates [Go modules](https://blog.friendsofgo.tech/posts/go-modules-en-tres-pasos/), which facilitates the download of the different packages/dependencies of the project.
```sh
go mod init     // Creates the file where dependencies are saved
go mod tidy     // Tidies and downloads the dependencies
```

##### API Usage

The request to the backend should always be a `POST`, and this could be an example JSON body for the request:
```json
{
    "destination": "Spain",
    "origin": "France"
}
```

The request **must have a 'X-Auth-Token' with the API-KEY** (for now the token is `SUPER_SECRET_API_KEY`, original, I know xD) if not, a `401 Unauthorized` code will be given.
To enforce the api key, a middleware is being used, which is added to the "router" so every time it receives a request the auth check is done.

The same endpoint is also implemented with `GET`, but not being used at the moment. 

When running the *Gin router* previously I used `run`, which serves HTTP. But, since being deployed, I use `runTLS`, which serves HTTPS. In this case you need to provide two certificates, later explained in [Domains and Cloudfare](#domains-and-cloudfare).

So far, the **response time in local is about 1ms**, while in AWS is around 36ms. In both cases has been stress tested with thousands of requests every 1ms, and has been able to not drop a single request.

To stress test the backend, I used the chrome extension named [RestfulStress](https://chrome.google.com/webstore/detail/restful-stress/lljgneahfmgjmpglpbhmkangancgdgeb).

##### CORS

Cross-Origin Resource Sharing (CORS) is a mechanism that uses additional HTTP headers to tell browsers to give a web application running at one origin, access to selected resources from a different origin. 
A web application executes a cross-origin HTTP request when it requests a resource that has a different origin (domain, protocol, or port) from its own.

In the backend, when responding to the requests there is the following headers that must be added to the response, so it **complies with the CORS policies**:
```golang
    c.Header("Access-Control-Allow-Origin", "*")  
    c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
```

"Preflighted" requests first send an HTTP request by the OPTIONS method to the resource on the other domain, to determine if the actual request is safe to send. 
Cross-site requests are preflighted like this since they may have implications to user data (my case, as **frontend and backend are hosted separately**).
In case the request is a CORS preflight (OPTIONS request), we will also, in case that we use an API key, add the following header ("X-Auth-Token"), so the client knows that the requests must contain an API key/token:
```golang
    c.Header("Access-Control-Allow-Origin", "*")  	
    c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers, X-Auth-Token")
```

The "Access-Control-Allow-Origin", determines what origin/website can use the endpoint. I could configure it, so the backend can only be used by `canitravelto.com` and my own IP. In that case, I should also include an [extra header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin): "Vary: Origin".

### Data Retriever
Coded in Go. Responsible for **updating the Covid daily data**, in the future will also handle other Database related functions.
It uses a Go-Cronjob to update the data every day at 10:30 AM.

### Business Handler
Coded in Go. Responsible for **handling all the API requests**. Uses Gin-gonic to handle the endpoints in **HTTPS mode**, so the content can be served to the HTTPS frontend.

## Docker
The different microservices are being run with **docker-compose in the EC2 AWS instance**. The images are hosted in a Github private docker registry for this project. I added an automation, so the older images are deleted once a month, or when a limit is reached.
There are two different Dockerfile's for each microservice, plus the docker-compose file to launch them together, plus create the "internal network", so they can communicate.

Performance wise, the difference between the compiled binary and the docker images has been negligible. Both are extremely fast, averaging around 53ms per request both. The backend itself, from when it receives the request till it sends back the response just takes 5 or 6ms.

Backend "logic" (Docker) response time, 5ms (POST/GET), 12micros (OPTIONS):
![Backend "logic" (Docker) response time](https://github.com/marcllort/CanITravelTo_Backend/blob/master/Documentation/Assets/backend-response-time.PNG)

Binary response time, 51ms:

![Binary response time](https://github.com/marcllort/CanITravelTo_Backend/blob/master/Documentation/Assets/binary-response-time.PNG)

Docker response time, 53ms:

![Docker response time](https://github.com/marcllort/CanITravelTo_Backend/blob/master/Documentation/Assets/docker-response-time.PNG)

### Kubernetes
Even though is **REALLY overkill for this project**, due to the small amount of visitors received, I wanted to implement Kubernetes to handle the docker containers.
I haven't been able to make it work in the production server (EC2 t2.micro instance) due to the small amount of resources. It makes the server unusable, always at 100% CPU and RAM usage.

I have a simple implementation of the project running in Kubernetes in my local environment, which I'll upload its configuration when working properly.

## Git
I'm using a **mono-repo**, as its enables me to share the docker-compose, readme, credentials... Later on, during the CI/CD is much easier to deal, as there is only one git repository to pull and deal with.
I also use the [Github Projects](https://github.com/marcllort/CanITravelTo_Backend/projects/1) feature, with the **Kanban methodology** to organize the new "stories" I have to develop/fix.

### CI/CD
I'm using **Github Actions**, to have everything centralized in Github. It uses a YML file, really similar to BitBucket/Gitlab or Jenkins.

So far I have [two pipelines](https://github.com/marcllort/CanITravelTo_Backend/actions), the CI ([ci.yml](https://github.com/marcllort/CanITravelTo_Backend/blob/master/.github/workflows/ci.yml)) and CD ([cd.yml](https://github.com/marcllort/CanITravelTo_Backend/blob/master/.github/workflows/cd.yml)). 

In the CI pipeline the steps implemented are: build the two microservices, run the unit and integration/E2E tests of each microservice. If it fails it will notify me through an email.

The Unit tests are done with the vanilla golang test methodology, similar to JUnit. The E2E tests, are a collection of Postman calls/tests that are being **run in the pipeline with Newman** (cli version of Postman).
The tests are written in Javascript, here there is a simple example:

```javascript
pm.test("Status code is 200", function () {
    pm.response.to.have.status(200);
});

pm.test("No error messages", function () {
    var jsonData = pm.response.json();
    pm.expect(jsonData.error).to.eql("");
});

pm.test("Origin/Destination Correct", function () {
    var jsonData = pm.response.json();
    pm.expect(jsonData.origin).to.eql("France");
    pm.expect(jsonData.destination).to.eql("Spain");
});
```

In the CD pipeline, if the commit is in the master branch, and it's a commit marked as a release, it will decrypt the Credentials needed to build each microservice image (explained in the next [section](#ssh-from-pipeline)), 
**build and upload the new Docker images** (to Github's docker registry), and then SSH into the EC2 instance to stop the old docker images and start the updated images with the latest changes.

Another solution, could be building the docker images in the server instead of doing it directly in the pipeline. This would allow me to just have the Credentials in my server, and avoid uploading the encrypted version to Github.
In case this was a long-term project, or a business it would probably be better to keep the Creds only on your server, but again, the Repository would be private, so the security shouldn't be a big problem... Make your choice!
I decided to use both solutions, the building of the image in the pipeline, and the deployment through an SSH script.

If we are in another branch (other than master) it will only run the CI pipeline (build and unit/integration tests).


#### SSH from pipeline
To ssh to AWS, a pem/key file is needed. As it would be really insecure to upload the key to github, I'm using a workaround (also used for the Credentials in the CI/CD pipelines). I encrypted (and uploaded to github) the PEM file using:
```sh
gpg --quiet --batch --yes --decrypt --passphrase="XXXX" --output key-aws.pem BusinessHandler/Creds/key-aws.pem.gpg
```

Then in the pipeline I decrypt the file using the passphrase (which is saved as a Secret environment variable), change the permissions and SSH (plus run the update script).

Important using the `-oStrictHostKeyChecking=no` when SSHing from a script/pipeline, so it automatically accepts the ECSDA key.

```sh
gpg --quiet --batch --yes --decrypt --passphrase="$LARGE_SECRET_PASSPHRASE" --output key-aws.pem BusinessHandler/Creds/key-aws.pem.gpg
chmod 600 key-aws.pem
ssh -oStrictHostKeyChecking=no -i key-aws.pem ubuntu@ec2-35-180-85-2.eu-west-3.compute.amazonaws.com "chmod +x CanITravelTo_Backend/Documentation/Ubuntu/update.sh && exit"
ssh -oStrictHostKeyChecking=no -i key-aws.pem ubuntu@ec2-35-180-85-2.eu-west-3.compute.amazonaws.com "CanITravelTo_Backend/Documentation/Ubuntu/update.sh $DB_PASSWORD && exit"
```


## Frontend

The frontend is a static website **coded in vanilla HTML/CSS/JS and hosted in Github Pages**, which is free with a maximum of a 100GB of bandwidth per month. To avoid this limitation, Cloudfare can be used. 
Cloudfare will (for free) cache the website in their servers and provide a Secure SSL certificate. To do so, follow [this](https://www.toptal.com/github/unlimited-scale-web-hosting-github-pages-cloudflare) tutorial.

In the future, I plan to move to a React frontend. Already have a React implementation running, but so far is not as nice as the vanilla one, because I don't have much experience with it! 
Once I'm done with CI/CD and backend tests, I'll continue with the React trainings to improve it.

## Domains and Cloudfare

Currently, two domains are being used: canitravelto.com (where the frontend is hosted, Github Pages) and canitravelto.wtf (where the backend API is hosted, EC2 AWS).

Both are using Cloudfare, which provides caching for the website, important for the frontend as Github offers a maximum of 100GB of bandwidth per month, but most importantly it provides TLS certificates, so the website and backend are HTTPS encrypted and safe to use.

At first, only the frontend used Cloudfare as there was no need for the backend to use HTTPS, until I saw that an HTTPS website can't consume from an HTTP API. The options where to go back to HTTP in the frontend (I didn't manage to do it, because Github Pages always provides HTTPS).

The second option was to serve the API as HTTPS. I created my own certificates, but as they were self-signed, HTTPS didn't like them, so I had to have valid certificates. To get an SSL certificate, you need a domain name, so I then acquired canitravelto.wtf, to use it instead of the public AWS IP.

[`canitravelto.com`](https://canitravelto.com) *is hosted by godaddy.com*

[`canitravelto.wtf`](https://canitravelto.wtf) *is hosted by name.com*

### Cloudfare
To configure Cloudfare in both domains, there's a few steps to follow:

#### Frontend
In the case of Github Pages ([canitravelto.com](https://canitravelto.com)) just follow the Cloudfare set-up. Once the email is received about your website being active, navigate to SSL/TLS and change the mode to FULL. If not done, the webpage won't be reached (still don't know why).

Another change to be done, is [CORS](#cors) in the API requests. [CORS](#cors) stands for Cross-origin resource sharing, which means that it will consume resources from another site. If not correctly configured, the requests will fail.

To make this work, in the frontend we only need to **add to the header of the request this api key**:

 `myHeaders.append("X-Auth-Token", "SUPER_SECRET_API_KEY")`


#### Backend
With the backend ([canitravelto.wtf](https://canitravelto.wtf)), first I hosted in [AWS with Route53](https://www.youtube.com/watch?v=qor31Egu0Rg) (probably not needed) and then did the same configuration as with Github pages in Cloudfare.

In this case we also have to navigate to SSL/TLS -> Origin Server and select Create Server. We need to copy the two keys and save them in `http-server.key` and `http-server.cert`. 

In the GoLang backend (`main.go`):

`router.RunTLS(PORT, "Creds/https-server.crt", "Creds/https-server.key")`

We also need to enable [CORS](#cors) on the backend, to **enable the requests from the frontend**. Just add this two headers to the returned JSON:

`c.Header("Access-Control-Allow-Origin", "*")`
    
`c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")`
