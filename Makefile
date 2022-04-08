BUILD_FOLDER = target
BUILD_OPTIONS =

OS := darwin
ARCH := amd64
BINARY_NAME := "dotup"
FULL_BINARY_NAME := $(BINARY_NAME)-$(OS)-$(ARCH)

PROJECT_USERNAME := lonepeon
PROJECT_REPOSITORY := dotup

GIT_BIN := git

GIT_COMMIT := $(shell $(GIT_BIN) rev-parse HEAD)
GIT_BRANCH := $(shell $(GIT_BIN) branch --no-color | awk '/^\* / { print $$2 }')
GIT_STATE := $(shell if [ -z "$(shell $(GIT_BIN) status --short)" ]; then echo clean; else echo dirty; fi)

GO_BIN := go
GO_FMT_BIN := gofmt
GO_STATICCHECK_BIN := $(GO_BIN) run ./vendor/honnef.co/go/tools/cmd/staticcheck

release: compile

compile:
	@echo "+$@"
	@touch internal/build/version.go
	@GOOS=$(OS) GOARCH=$(ARCH) go build $(BUILD_OPTIONS) \
		-ldflags \
			 "-X github.com/lonepeon/dotup/internal/build.gitBranch=$(GIT_BRANCH) \
		 	  -X github.com/lonepeon/dotup/internal/build.gitCommit=$(GIT_COMMIT) \
			  -X github.com/lonepeon/dotup/internal/build.gitState=$(GIT_STATE) \
			  -X github.com/lonepeon/dotup/internal/build.buildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')" \
		-o $(BUILD_FOLDER)/$(FULL_BINARY_NAME)

test: test-unit test-format test-security

test-unit:
	@echo "+$@"
	@go test ./...

test-format:
	@echo "+ $@"
	@test -z "$$($(GO_BIN) fmt ./... | tee /dev/stderr)" || \
	  ( >&2 echo "=> please format Go code with '$(GO_BIN) fmt ./...'" && false)

test-security:
	@echo "+ $@"
	@$(GO_STATICCHECK_BIN) ./...
