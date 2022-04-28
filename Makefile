release-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./build/forge-stats-linux-amd64

release-linux-arm64:
	GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o ./build/forge-stats-linux-arm64

release: release-linux-amd64 release-linux-arm64

debug:
	go build
