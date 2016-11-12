strato:
	CGO_ENABLED=0 go build -tags netgo -a -v

image: strato
	docker build -t strato .

run:
	docker run -it strato sh
