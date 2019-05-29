BINARY_NAME=hellonow

build:
		go build -o $(BINARY_NAME) ./app

run:
		go run ./...

clean:
		rm -f $(BINARY_NAME)
