
PROJECT=parun
BIN=$(CURDIR)/bin
EXEC=$(PROJECT).exe


all: build test

build:
	go build -o $(BIN)/$(EXEC)

test:
	go test -v 

dep:
	go mod tidy
	
cc:
	SET GOOS=linux& SET GOARCH=amd64& go build -o $(BIN)/linux/$(EXEC) 
	SET GOOS=windows& SET GOARCH=amd64& go build -o $(BIN)/windows/$(EXEC) 