lay:
	CGO_ENABLED=0 go build -tags netgo -a -v

image: lay
	docker build -t lay .

run:
	docker run -it lay sh
