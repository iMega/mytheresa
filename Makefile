REPO = github.com/imega/mytheresa
CWD = /go/src/$(REPO)
GO_IMG = golang:alpine

test: unit

unit:
	@docker run --rm -v $(CURDIR):$(CWD) \
		-w $(CWD) \
		-e GOFLAGS=-mod=mod \
		$(GO_IMG) \
		sh -c 'go test -v `go list ./...`'
