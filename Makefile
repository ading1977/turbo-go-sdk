SOURCE_DIRS = pkg
PACKAGES := go list ./... | grep -v /vendor | grep -v /out

build:
	go build ./...

test:
	@go test -v -race ./pkg/...

bin=utility
utility: clean
	@go build -o ${bin} ./cmd

.PHONY: fmtcheck
fmtcheck:
	@gofmt -l $(SOURCE_DIRS) | grep ".*\.go"; if [ "$$?" = "0" ]; then exit 1; fi

.PHONY: vet
 vet:
	@go vet $(shell $(PACKAGES))

clean:
	@rm -rf ${bin}
