.PHONY: proto proto_format

proto: proto_format
	protoc \
	-I=./ \
	-I=${GOPATH}/src \
	--go_out=./dist \
	--go_opt=module=github.com/jun06t/bazel-sample/protobuf/proto \
	./proto/*.proto

proto_format:
	find ./proto/ -name "*.proto" | xargs clang-format -i

