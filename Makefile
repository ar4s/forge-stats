release:
	go build -ldflags "-s -w"

debug:
	go build
