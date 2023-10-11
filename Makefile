build:
	go build -o bin/go-sort

run: build
	./bin/go-sort

test:
	go test -v ./...