NAME = golang-template

MAIN := ./cmd/
BINARY := bin/$(NAME)

Version := $(shell cat Version)
COMMIT := $(shell git rev-parse --short HEAD)
Branch := $(shell git rev-parse --abbrev-ref HEAD)
Builder := $(shell whoami)
BuildAt := $(shell date +%F.%T)

CTIMEVAR = -X main.Commit=$(COMMIT) \
        -X main.Version=$(Version) \
        -X main.Branch=$(Branch) \
        -X main.Builder=$(Builder) \
        -X main.BuildAt=$(BuildAt)
GO_LDFLAGS = -ldflags "-w $(CTIMEVAR) -s"


.PHONY: clean
clean:
	rm -rf data/*
	rm -rf bin/*

.PHONY: build
build:
	go build -tags "$(BUILDTAGS)" $(GO_LDFLAGS) -o $(BINARY) $(MAIN)

.PHONY: test
test:
	go test -v --cover ./...

.PHONY: dev
dev: clean asset build
	ENV_DEBUG=true ENV_SERVER_PPROF=true $(BINARY)

.PHONY: pprof
pprof: clean asset build
	ENV_SERVER_PPROF=true $(BINARY)


.PHONY: docker-image docker-base push-build-base
docker-image:
	docker build \
		-t wrfly/$(NAME) \
		-t wrfly/$(NAME):develop \
		-t wrfly/$(NAME):$(VERSION) \
		-f docker/Dockerfile .

docker-base:
	docker build \
		-t wrfly/$(NAME):build-base \
		-f docker/Dockerfile.base .

push-build-base:
	docker push wrfly/$(NAME):build-base

.PHONY: push-img
push-img:
	docker push wrfly/$(NAME)
	docker push wrfly/$(NAME):$(VERSION)

.PHONY: push-dev-img
push-dev-img:
	docker push wrfly/$(NAME):develop

.PHONY: tools
tools:
	go get github.com/wrfly/bindata

.PHONY: asset
asset:
	bindata -pkg asset \
		-resource html \
		-target asset
