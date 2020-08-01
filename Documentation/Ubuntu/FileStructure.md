.
├── CanITravelTo_Backend
│   ├── BusinessHandler
│   │   ├── Controller
│   │   │   ├── CovidController.go
│   │   │   ├── GetHandler.go
│   │   │   ├── OptionsHandler.go
│   │   │   ├── PostHandler.go
│   │   │   └── ResponseHandler.go
│   │   ├── Creds
│   │   │   ├── creds.json
│   │   │   ├── https-server.crt
│   │   │   └── https-server.key
│   │   ├── Database
│   │   │   ├── mysqlConnection.go
│   │   │   ├── travel_Covid.sql
│   │   │   └── travel_Passport.sql
│   │   ├── DockerCommands
│   │   ├── Dockerfile
│   │   ├── Middleware
│   │   │   └── WhiteList.go
│   │   ├── Model
│   │   │   ├── APIRequest.go
│   │   │   ├── Covid.go
│   │   │   └── InfoCountry.go
│   │   ├── Scripts
│   │   │   ├── CanITravelToGlobal.postman_collection.json
│   │   │   ├── CanITravelToLocal.postman_collection.json
│   │   │   └── updateServer.sh
│   │   ├── Utils
│   │   │   └── Credentials.go
│   │   ├── go.mod
│   │   ├── go.sum
│   │   └── main.go
│   ├── DataRetriever
│   │   ├── Controller
│   │   │   ├── CovidController.go
│   │   │   └── DatabaseHandler.go
│   │   ├── Creds
│   │   │   ├── creds.json
│   │   │   ├── https-server.crt
│   │   │   └── https-server.key
│   │   ├── Database
│   │   │   ├── mysqlConnection.go
│   │   │   ├── travel_Covid.sql
│   │   │   └── travel_Passport.sql
│   │   ├── Dockerfile
│   │   ├── Model
│   │   │   ├── Covid.go
│   │   │   └── InfoCountry.go
│   │   ├── Utils
│   │   │   └── Credentials.go
│   │   ├── go.mod
│   │   ├── go.sum
│   │   └── main.go
│   ├── docker-compose.yaml
│   ├── readme.md
│   └── readme_deprecated.md
├── OldBinary ----------> IMPORTANT
│   ├── CanITravelTo
│   ├── Creds ----------> IMPORTANT
│   │   ├── creds.json
│   │   ├── https-server.crt
│   │   └── https-server.key
│   ├── server.log
│   └── updateServer.sh
└── update.sh

17 directories, 49 files