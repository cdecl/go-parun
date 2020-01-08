
GOPATH=$(CURDIR)
GOBIN=$(GOPATH)/bin
GOFILES=parun
EXEC=parun.exe


all: build test

build:
	@env GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o $(GOBIN)/$(EXEC) $(GOFILES)

run:
	@env GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run $(GOFILES)

test:
	@env GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -v $(GOFILES)

get:
	-env GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get -d ./...
	
cc:
	@env GOPATH=$(GOPATH) GOBIN=$(GOBIN)  GOOS=linux GOARCH=amd64 go build -o $(GOBIN)/linux/$(EXEC) $(GOFILES)
	@env GOPATH=$(GOPATH) GOBIN=$(GOBIN)  GOOS=windows GOARCH=amd64 go build -o $(GOBIN)/windows/$(EXEC) $(GOFILES)