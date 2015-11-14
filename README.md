# go-protobuf-example

```
$ brew install --devel protobuf
$ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
```
```
$ protoc -I ./protos ./protos/helloworld.proto --go_out=plugins=grpc:helloworld
```
```
$ go run server/main.go
$ go run client/main.go
```
