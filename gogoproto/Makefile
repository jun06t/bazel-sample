.PHONY: proto proto_format

proto: proto_format
	protoc \
	-I=./ \
	-I=${GOPATH}/src \
	--go_out=plugins=grpc:./dist \
	--go_opt=module=github.com/jun06t/bazel-sample/gogoproto/proto \
	./proto/*.proto

gogoproto: proto_format
	protoc \
	-I=./ \
	-I=${GOPATH}/src \
	--gogo_out=${GOPATH}/src \
	./proto/*.proto

proto_format:
	find ./proto/ -name "*.proto" | xargs clang-format -i

