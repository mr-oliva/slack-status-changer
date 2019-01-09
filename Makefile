.PHONY: build
build:
	GOOS=windows GOARCH=amd64 go build -o slack-status-changer.exe main.go
	GOOS=linux GOARCH=amd64 go build -o slack-status-changer_linux main.go
	GOOS=darwin GOARCH=amd64 go build -o slack-status-changer_osx main.go
