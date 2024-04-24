project-name=gui-fresheye
build_filename=./build/$(project-name)
go_ldflags="-w -s"

# Выполнить проверку линтером
# go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2
# doc: https://golangci-lint.run/usage/install/
.PHONY: lint
lint:
	golangci-lint run ./...
