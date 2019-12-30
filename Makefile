
GOPATH=$(CURDIR)
GOBIN=$(GOPATH)/bin
GOFILES=./...
EXEC=parun.exe

build:
	@env GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o $(GOBIN)/$(EXEC) $(GOFILES)

run:
	@env GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run $(GOFILES)

test:
	@env GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test $(GOFILES)

get:
	# -env GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get -u github.com/shirou/gopsutil
	-env GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get -u ./...
	
cc:
	@env GOPATH=$(GOPATH) GOBIN=$(GOBIN)  GOOS=linux GOARCH=amd64 go build -o $(GOBIN)/linux/$(EXEC) $(GOFILES)
	@env GOPATH=$(GOPATH) GOBIN=$(GOBIN)  GOOS=windows GOARCH=amd64 go build -o $(GOBIN)/windows/$(EXEC) $(GOFILES)