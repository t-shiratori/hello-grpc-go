# hello-grpc-go

Hello World for gRPC with Go

## Download Packages

```sh
go mod tidy
```

## Run Server

```sh
go run server/main.go
server is runnig...
```

## Run Client

```sh
% go run client/main.go
reqWords: 
Please need words 'Pin'!
```

```sh
% go run client/main.go Pin
reqWords: Pin
Pon!
```
