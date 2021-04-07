.PHONY: build/ensureGoPath
build/ensureGoPath:
	@$(shell if [ ! -d go ]; then mkdir go;fi)

ci/binary: build/ensureGoPath
ci/binary: export GOPATH=$(abspath go)
ci/binary: export GO111MODULE=on
ci/binary: export GOPROXY=https://goproxy.cn
ci/binary: export GOOS=linux
ci/binary: export CGO_ENABLED=1

clean:
	@rm -f $(NAME)

fmt:
	command -v gofumpt || (WORK=$(shell pwd) && cd /tmp && GO111MODULE=on go get mvdan.cc/gofumpt && cd $(WORK))
	gofumpt -w -s -d .
	go vet "./..."

lint:
	golangci-lint run  -v

ci/lint: export GOPATH=$(abspath go)
ci/lint: export GO111MODULE=on
ci/lint: export GOPROXY=https://goproxy.cn
ci/lint: export GOOS=linux
ci/lint: export CGO_ENABLED=1
ci/lint: export GOPATH=$(abspath go)
ci/lint: build/ensureGoPath lint

