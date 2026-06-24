APP_NAME=video-renamer

build:
	go build -o bin/$(APP_NAME) ./cmd

windows:
	GOOS=windows GOARCH=amd64 go build -o bin/$(APP_NAME).exe ./cmd

linux:
	GOOS=linux GOARCH=amd64 go build -o bin/$(APP_NAME)-linux ./cmd

mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/$(APP_NAME)-mac ./cmd

release:
	mkdir -p bin
	GOOS=windows GOARCH=amd64 go build -o bin/$(APP_NAME).exe ./cmd
	GOOS=linux GOARCH=amd64 go build -o bin/$(APP_NAME)-linux ./cmd
	GOOS=darwin GOARCH=amd64 go build -o bin/$(APP_NAME)-mac ./cmd
	-ldflags "-X renamer/internal/version.Version=v1.0.0"

test:
	go test ./...

clean:
	rm -rf bin