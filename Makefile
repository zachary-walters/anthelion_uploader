build_windows:
	GOOS=windows GOARCH=amd64 go build -o bin/anthelion_uploader.exe main.go

build_windows_32:
	GOOS=windows GOARCH=386 go build -o bin/anthelion_uploader.exe main.go

build_mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/anthelion_uploader main.go

build_mac_32:
	GOOS=darwin GOARCH=386 go build -o bin/anthelion_uploader main.go

build_linux:
	GOOS=linux GOARCH=amd64 go build -o bin/anthelion_uploader main.go

build_linux_32:
	GOOS=linux GOARCH=386 go build -o bin/anthelion_uploader main.go
