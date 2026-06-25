APP_NAME=video-renamer
VERSION=v1.0.0
LDFLAGS=-ldflags "-X renamer/internal/ui.Version=$(VERSION)"

ifeq ($(OS),Windows_NT)

EXE=.exe

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
	go build $(LDFLAGS) -o bin/$(APP_NAME)$(EXE) ./cmd

windows:
	go build $(LDFLAGS) -o bin/$(APP_NAME).exe ./cmd

linux:
	cmd /C "set GOOS=linux&& set GOARCH=amd64&& go build $(LDFLAGS) -o bin/$(APP_NAME)-linux ./cmd"

mac:
	cmd /C "set GOOS=darwin&& set GOARCH=amd64&& go build $(LDFLAGS) -o bin/$(APP_NAME)-mac ./cmd"

<<<<<<< HEAD
release: windows linux mac

clean:
	if exist bin rmdir /S /Q bin

else

build:
	go build $(LDFLAGS) -o bin/$(APP_NAME) ./cmd

windows:
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o bin/$(APP_NAME).exe ./cmd

linux:
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/$(APP_NAME)-linux ./cmd

mac:
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o bin/$(APP_NAME)-mac ./cmd

release: windows linux mac

=======
>>>>>>> 57291899d03484055e61b5f3615c498533558ae2
clean:
	rm -rf bin

endif

<<<<<<< HEAD
test:
	go test ./...
=======
release:
	mkdir -p bin
	go build -o bin/$(APP_NAME) ./cmd

test:
	go test ./...
>>>>>>> 57291899d03484055e61b5f3615c498533558ae2
