all:
	@echo "Make targets:"
	@echo "deps: fetch all dependencies required"
	@echo "devel: run the development server using gin"
	@echo "run: run the development server without using gin"

deps:
	go get -u -v github.com/codegangsta/gin/...
	$(GOPATH)/bin/godeps -u dependencies.tsv

devel:
	$(GOPATH)/bin/gin run config/development.yaml

run:
	go run server.go config/development.yaml

.PHONY: all deps devel run
