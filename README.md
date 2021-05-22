# GRPC Hello World

## Requirements
* MacOs
```sh
brew install protobuf
```
* Linux
```sh
$ curl -OL https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip
$ unzip protoc-3.5.1-linux-x86_64.zip -d protoc3
$ sudo mv protoc3/bin/* /usr/local/bin/
$ sudo mv protoc3/include/* /usr/local/include/

$ sudo chown [user] /usr/local/bin/protoc
$ sudo chown -R [user] /usr/local/include/google
```

## Commands
**Compile proto file**
```sh
$ protoc hello/hello.proto --go_out=plugins=grpc:.
```

## Run

* Server
```sh
$ gor hello/server/server.go
```

* Client
```sh
$ gor hello/client/client.go
```
