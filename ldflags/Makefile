
PACKAGE_PATH := github.com/jun06t/bazel-sample/ldflags/alive
VERSION ?= $(shell git rev-parse --verify HEAD)
DATE ?= $(shell date '+%Y/%m/%d %H:%M:%S %Z')

build:
	go build -ldflags "-X $(PACKAGE_PATH).BuildVersion=$(VERSION) -X \"$(PACKAGE_PATH).Timestamp=$(DATE)\"" cmd/main.go

bazel-run:
	BUILDABLE=ldflags bazel run --stamp --workspace_status_command=$$(pwd)/tag.sh //cmd
