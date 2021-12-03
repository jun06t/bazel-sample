# gRPC Bazel sample

## Compile
```
bazel build //proto
cp bazel-bin/proto/helloworld_go_proto_/github.com/jun06t/bazel-sample/grpc/proto/helloworld.pb.go proto/
```

## Run gRPC server
```
bazel run //server
```

## Run gRRC Client
```
bazel run //client
```

Then it returns
```
2021/12/03 18:28:08 Reply:  Hello alice
```
