strato:
	go build -ldflags="-s -w"

base: strato
	./scripts/build-base

run: base
	docker run -it strato sh
