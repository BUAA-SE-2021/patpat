build:
	env GOOS=windows GOARCH=amd64 go build -o bin/patpat-windows-amd64.exe main.go
	env GOOS=linux GOARCH=amd64 go build -o bin/patpat-linux-amd64 main.go
	env GOOS=darwin GOARCH=amd64 go build -o bin/patpat-macos-amd64 main.go