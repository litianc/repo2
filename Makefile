BINARY_NAME=hello
build:
	go build -o $(BINARY_NAME) -v
.PHONY: build

run:
	go build -o $(BINARY_NAME) -v
	./$(BINARY_NAME)
.PHONY: run
