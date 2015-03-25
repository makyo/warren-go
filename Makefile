ifndef GOPATH
	$(warning You need to set up a GOPATH)
endif

PROJECT := github.com/makyo/warren-go
PROJECT_DIR := $(shell go list -e -f '{{.Dir}}' $(PROJECT))

help:
	@echo "Available targets:"
	@echo "  deps - fetch all dependencies required"
	@echo "  devel - run the development server using gin"
	@echo "  run - run the development server without using gin"
	@echo "  create-deps - rebuild the dependencies.tsv file"
	@echo "  build - build the application"
	@echo "  check - run tests"
	@echo "  install - Install the application in your GOPATH"
	@echo "  clean - clean the project"

deps: $(GOPATH)/bin/godeps
	go get -v github.com/codegangsta/gin/...
	$(GOPATH)/bin/godeps -u dependencies.tsv

create-deps: $(GOPATH)/bin/godeps
	godeps -t $(shell go list $(PROJECT)/...) > dependencies.tsv || true

ifeq ($(CURDIR),$(PROJECT_DIR))

devel:
	$(GOPATH)/bin/gin run config/development.yaml

run:
	go run server.go config/development.yaml

build:
	go build $(PROJECT)/...

check:
	go test $(PROJECT)/...

clean:
	go clean $(PROJECT)/...

install:
	go install $(INSTALL_FLAGS) -v $(PROJECT)/...

else

devel:
	$(error Cannot $@; $(CURDIR) is not on GOPATH)

run:
	$(error Cannot $@; $(CURDIR) is not on GOPATH)

build:
	$(error Cannot $@; $(CURDIR) is not on GOPATH)

check:
	$(error Cannot $@; $(CURDIR) is not on GOPATH)

install:
	$(error Cannot $@; $(CURDIR) is not on GOPATH)

clean:
	$(error Cannot $@; $(CURDIR) is not on GOPATH)

endif

.PHONY: all build check clean create-deps deps devel install run
