BAZEL := bazelisk

TESTABLE := $(shell ${BAZEL} query 'kind("go_test", //...)')

.PHONY: test
test:
	${BAZEL} test -- ${TESTABLE}

.PHONY: test-race
test-race:
	${BAZEL} test --@io_bazel_rules_go//go/config:race -- ${TESTABLE}

.PHONY: test-integration
test-integration:
	${BAZEL} test --@io_bazel_rules_go//go/config:tags=integration -- ${TESTABLE}

coverage:
	${BAZEL} coverage -- ${TESTABLE}
	echo "mode: set" > coverage.out
	find bazel-out/ -name "coverage.dat" | xargs tail -q -n +2 >> coverage.out

.PHONY: gazelle
gazelle:
	${BAZEL} run //:gazelle -- update -build_tags integration

.PHONY: gazelle-update-repos
gazelle-update-repos:
	${BAZEL} run //:gazelle-update-repos

