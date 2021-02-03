build:
	env GOOS=windows GOARCH=amd64 go build -o bin/patpat-windows-64.exe main.go
	env GOOS=linux GOARCH=amd64 go build -o bin/patpat-linux-64 main.go
	env GOOS=darwin GOARCH=amd64 go build -o bin/patpat-macos-64 main.go