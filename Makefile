strato:
	go build -ldflags="-s -w"

base: strato
	docker build -t strato .

run: base
	docker run -it strato sh
