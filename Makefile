export GO111MODULE=on

help:
	@echo "This is a helper makefile for oapi-codegen"
	@echo "Targets:"
	@echo "    setup:              install the development dependencies"
	@echo "    generate:           rerun code generation"
	@echo "    lint:               run golangci-lint"
	@echo "    test-unit:          run unit tests"
	@echo "    report-test-unit:   run unit tests and generate report file"

setup:
	# install golanglint-ci into ./bin
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.16.0
	# Install gobin globally
	env GO111MODULE=off go get -u github.com/myitcv/gobin

generate:
	go generate ./...

lint:
	./bin/golangci-lint run ./...

test: test-unit test-integration

report-test: report-test-unit

test-unit:
	go test -count=1 ./cmd/... ./internal/...

report-test-unit:
	go test -count=1 -coverprofile=reports/test-unit.out -v -p 5 ./cmd/... ./internal/... | gobin -m -run github.com/apg/patter > reports/test-unit.tap

.PHONY: \
	setup \
	generate \
	lint \
	test-unit \
	report-test-unit \
