export PATH=$PATH:$(go env GOPATH)/bin
BINARY_NAME=server
SRC_PATH=./

build:
	go build -o $(BINARY_NAME) $(SRC_PATH)

run: build
	./$(BINARY_NAME)

watch:
	reflex -sr go make run
