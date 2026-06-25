APP_NAME=video-renamer

ifeq ($(OS),Windows_NT)
EXE=.exe

build:
	go build -o bin/$(APP_NAME)$(EXE) ./cmd

windows:
	go build -o bin/$(APP_NAME).exe ./cmd

linux:
	cmd /C "set GOOS=linux&& set GOARCH=amd64&& go build -o bin/$(APP_NAME)-linux ./cmd"

mac:
	cmd /C "set GOOS=darwin&& set GOARCH=amd64&& go build -o bin/$(APP_NAME)-mac ./cmd"

clean:
	if exist bin rmdir /S /Q bin

else

build:
	go build -o bin/$(APP_NAME) ./cmd

windows:
	GOOS=windows GOARCH=amd64 go build -o bin/$(APP_NAME).exe ./cmd

linux:
	GOOS=linux GOARCH=amd64 go build -o bin/$(APP_NAME)-linux ./cmd

mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/$(APP_NAME)-mac ./cmd

clean:
	rm -rf bin

endif

release:
	mkdir -p bin
	go build -o bin/$(APP_NAME) ./cmd

test:
	go test ./...
