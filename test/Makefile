BAZEL := bazelisk
BAZEL_FLAGS =
BAZEL_COMMAND_FLAGS =

ifdef CI
	BAZEL_COMMAND_FLAGS += --config=ci
endif

TESTABLE := $(shell ${BAZEL} ${BAZEL_FLAGS} query 'kind("go_test", //go/...)')

.PHONY: test
test:
	${BAZEL} ${BAZEL_FLAGS} test ${BAZEL_COMMAND_FLAGS} -- ${TESTABLE}

.PHONY: test-race
test-race:
	${BAZEL} ${BAZEL_FLAGS} test ${BAZEL_COMMAND_FLAGS} --@io_bazel_rules_go//go/config:race -- ${TESTABLE}

.PHONY: test-integration
test-integration:
	${BAZEL} ${BAZEL_FLAGS} test ${BAZEL_COMMAND_FLAGS} --@io_bazel_rules_go//go/config:tags=integration -- ${TESTABLE}

.PHONY: gazelle
gazelle:
	${GO_ENV} ${BAZEL} ${BAZEL_FLAGS} run //:gazelle -- update -build_tags integration

.PHONY: gazelle-update-repos
gazelle-update-repos:
	${GO_ENV} ${BAZEL} ${BAZEL_FLAGS} run //:gazelle-update-repos

