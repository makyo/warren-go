all:
	@echo "Make targets:"
	@echo "deps: fetch all dependencies required"
	@echo "devel: run the development server using gin"

deps:
	go get -u -v github.com/codegangsta/gin/...
	$(GOPATH)/bin/godeps -u dependencies.tsv

devel:
	$(GOPATH)/bin/gin run config/development.yaml

.PHONY: all deps devel
