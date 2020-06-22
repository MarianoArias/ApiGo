# ApiGo


## Requirements 
* GO >= 1.9
* MySQL >= 8.0


## Installation 
* Clone repository
```
git clone git@github.com:MarianoArias/ApiGo.git
```


## Test 
* Run following command in the project's root folder:
```
APP_ENV=test go test ./cmd/customers -count=1 -coverprofile=cover.out
```
* Or run following command in the project's root folder for full testing:
```
APP_ENV=test go test ./... -count=1 -coverprofile=cover.out
```


## Parse and generate API documentation
* Run following command in the project's root folder:
```
$GOPATH/bin/swag init --generalInfo=cmd/customers/main.go --output=cmd/customers/docs
```


## Deploy
* Run following command in the project's root folder:
```
go run cmd/customers/main.go
```