REPO = github.com/imega/mytheresa
CWD = /go/src/$(REPO)
GO_IMG = golang:alpine

test: lint unit

unit:
	@docker run --rm -v $(CURDIR):$(CWD) \
		-w $(CWD) \
		-e GOFLAGS=-mod=mod \
		-e CGO_ENABLED=0 \
		$(GO_IMG) \
		sh -c 'go test -v `go list ./... | grep -v tests`'

lint:
	@docker run --rm -t -v $(CURDIR):$(CWD) -w $(CWD) \
		golangci/golangci-lint golangci-lint run

acceptance: clean
	@GO_IMG=$(GO_IMG) CWD=$(CWD) docker-compose up -d --build --scale acceptance=0
	@GO_IMG=$(GO_IMG) CWD=$(CWD) docker-compose up --abort-on-container-exit acceptance

clean:
	@GO_IMG=$(GO_IMG) CWD=$(CWD) docker-compose rm -sfv
