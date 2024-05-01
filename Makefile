project-name=gui-fresheye
build_filename=./build/$(project-name)
go_ldflags="-w -s"
app_id="io.fyne.fresheye"
app_version="0.0.1"

# Выполнить проверку линтером
# go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.1
# doc: https://golangci-lint.run/usage/install/
.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: generate
generate:
	go generate ./...

# go install github.com/fyne-io/fyne-cross@lates
.PHONY: build-windows
build-windows:
	fyne-cross windows -arch=amd64,386 -app-id $(app_id) -app-version $(app_version) .

.PHONY: build-linux
build-linux:
	fyne-cross linux -app-id $(app_id) -app-version $(app_version) .

.PHONY: build-darwin
	fyne-cross darwin -app-id $(app_id) -app-version $(app_version) .

.PHONY: clean
clean:
	rm -rf fyne-cross/ || true

.PHONY: release
release: generate lint build-windows build-linux build-darwin
