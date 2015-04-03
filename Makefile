PROJECT := github.com/warren-community/warren
PROJECT_DIR := $(shell go list -e -f '{{.Dir}}' $(PROJECT))
GODEPS := github.com/codegangsta/gin/...\
	github.com/smartystreets/goconvey/...\
	launchpad.net/godeps/...


NODE_TARGETS=node_modules/coffee_script

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

$(NODE_TARGETS): package.json
	npm install

coffee: $(NODE_TARGETS)
	node_modules/coffee-script/bin/coffee -o public/js -cw public/coffee

ifeq ($(CURDIR),$(PROJECT_DIR))

deps:
	go get -v $(GODEPS)
	godeps -u dependencies.tsv

create-deps:
	godeps -t $(shell go list $(PROJECT)/...) > dependencies.tsv || true

devel:
	${MAKE} -j2 coffee gin

gin:
	$(GOPATH)/bin/gin run config/development.yaml

run:
	go run server.go config/development.yaml

build:
	go build $(PROJECT)/...

check:
	go test -v $(PROJECT)/...

clean:
	go clean $(PROJECT)/...

install:
	go install $(INSTALL_FLAGS) -v $(PROJECT)/...

else

deps:
	$(error Cannot $@; $(CURDIR) is not on GOPATH)

create-deps:
	$(error Cannot $@; $(CURDIR) is not on GOPATH)

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

.PHONY: all build check clean coffee create-deps deps devel gin install run
