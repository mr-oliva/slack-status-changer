.PHONY: build
build:
	GOOS=windows GOARCH=amd64 go build -o slack-status-changer.exe cmd/slack-status-changer/main.go
	GOOS=linux GOARCH=amd64 go build -o slack-status-changer_linux cmd/slack-status-changer/main.go
	GOOS=darwin GOARCH=amd64 go build -o slack-status-changer_osx cmd/slack-status-changer/main.go
