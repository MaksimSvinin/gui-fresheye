project-name=gui-fresheye
build_filename=./build/$(project-name)
go_ldflags="-w -s"
app_id="io.fyne.fresheye"
app_version="0.1.0"

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
	fyne-cross linux -arch=amd64,386 -app-id $(app_id) -app-version $(app_version) .

.PHONY: release-build-windows
release-build-windows:
	fyne-cross windows -arch=amd64,386 -app-id $(app_id) -app-version $(app_version) -release .

.PHONY: release-build-linux
release-build-linux:
	fyne-cross linux -arch=amd64,386 -app-id $(app_id) -app-version $(app_version) -release .

.PHONY: clean
clean:
	rm -rf fyne-cross/ || true

.PHONY: rename-build-out
rename-build-out:
	mv fyne-cross/bin/linux-386/gui-fresheye fyne-cross/bin/linux-386/gui-fresheye-linux-386
	mv fyne-cross/bin/linux-amd64/gui-fresheye fyne-cross/bin/linux-amd64/gui-fresheye-linux-amd64
	mv fyne-cross/bin/windows-386/gui-fresheye.exe fyne-cross/bin/windows-386/gui-fresheye-windows-386.exe
	mv fyne-cross/bin/windows-amd64/gui-fresheye.exe fyne-cross/bin/windows-amd64/gui-fresheye-windows-amd64.exe

.PHONY: build
build: clean generate lint build-windows build-linux rename-build-out
