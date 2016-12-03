strato:
	CGO_ENABLED=0 go build -tags netgo -ldflags="-s -w" -a -v
