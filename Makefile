#!/usr/bin/make -f

export CGO_ENABLED=0

# Builds the project.
build:
	gox -os='linux darwin' -arch='amd64' -output='bin/skpr-crond_{{.OS}}_{{.Arch}}' -ldflags='-extldflags "-static"' github.com/skpr/crond/cmd/skpr-crond

# Run all lint checking with exit codes for CI.
lint:
	golint -set_exit_status `go list ./... | grep -v /vendor/`

# Run tests with coverage reporting.
test:
	go test -cover ./...

.PHONY: *
