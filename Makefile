IMG ?= relay:latest
TS := $(shell /bin/date "+%Y%m%d%H%M%S")
DEV_USER ?= paralusio
DEV_TAG := docker.io/${DEV_USER}/relay:$(TS)

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: vendor
vendor:
	GOPROXY=direct go mod vendor

check:
	go fmt ./...
	go vet ./...

build:
	docker build . -t ${IMG} --build-arg BUILD_USR=${BUILD_USER} --build-arg BUILD_PWD=${BUILD_PASSWORD}	

build-agent:
	docker build -t ${IMG} --build-arg BUILD_USR=${BUILD_USER} --build-arg BUILD_PWD=${BUILD_PASSWORD} -f Dockerfile.agent .

tag-dev:
	docker tag ${IMG} $(DEV_TAG)
	docker push $(DEV_TAG)

build-dev:
	rm -rf relay.*
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn" -o relay.dev main.go
	upx -5 -o relay.upx relay.dev	
	docker build -f Dockerfile.dev -t ${IMG} .
	$(MAKE) tag-dev

## changelog: generate changelog (latest release)
.PHONY: changelog
changelog:
	conventional-changelog -i CHANGELOG.md -s -p conventionalcommits
