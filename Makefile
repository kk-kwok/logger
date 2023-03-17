
ci/binary: export GOPATH=$(abspath go)
ci/binary: export GO111MODULE=on
ci/binary: export GOPROXY=https://goproxy.cn
ci/binary: export GOPRIVATE=code.hellotalk.com
ci/binary: export GOOS=linux
ci/binary: export CGO_ENABLED=1

fmt:
	command -v gofumpt || (WORK=$(shell pwd) && cd /tmp && GO111MODULE=on go get mvdan.cc/gofumpt && cd $(WORK))
	gofumpt -w -s -d .
	go vet "./..."

lint:
	golangci-lint run  -v

ci/lint: export GO111MODULE=on
ci/lint: export GOPROXY=https://goproxy.cn
ci/lint: export GOPRIVATE=code.hellotalk.com
ci/lint: export GOOS=linux
ci/lint: export CGO_ENABLED=0
ci/lint: lint

test:
	go test -v .

bench:
	go test -test.count 20 -test.benchmem -test.bench . --run=Benchmark_With

update:
	@curl -Ls https://code.hellotalk.com/snippets/11/raw > Makefile
	#@curl -Ls https://code.hellotalk.com/snippets/2/raw > .gitlab-ci.yml
	#@curl -Ls https://code.hellotalk.com/snippets/23/raw > .golangci.yml
