
# Smart Alarm - Internet of things - Project

## Backend

This is repository will include web server to store statistics from alarm gathering informations about sleep.

   Backend: Golang + MongoDB  
   Frontend: Angular2  

### Set up

Install [golang](https://golang.org/dl/)  
Install [mongoDB](https://www.mongodb.com/download-center?jmp=nav#community)  

Clone this repository  
  
Start mongodb
```
$ sudo service mongod start
```
Create database `iot`
```
$ mongo
> use iot
```
Now your database `iot` should be running on default port  

Start server
```
$ go get github.com/knopt/iot  
$ cd $GOPATH/src/github.com/knopt/iot/backend  
$ go get .
$ go install
$ $GOPATH/bin/backend
```

## Frontend
```
$ cd $GOPATH/src/github.com/knopt/iot/frontend
$ npm install
$ npm start
```
## Physical device software
TBA
