PKG := github.com/pfolta/cdrdao2audio

# Build metadata injected into the binary via -ldflags.
BUILD_DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
VERSION := $(shell git describe --match "v[0-9]*" --dirty="-m" --always --tags || echo "dev")

BUILD_DIR ?= ./build

# Use the requested target platform if provided.
# Otherwise, default to the platform of the build host.
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

ifeq ($(GOOS),windows)
	BINARY := cdrdao2audio.exe
else
	BINARY := cdrdao2audio
endif

GO_LDFLAGS := \
	-s \
	-w \
	-X "$(PKG).buildDate=$(BUILD_DATE)" \
	-X "$(PKG).version=$(VERSION)"

# Control ANSI color output.
#   - never:  Disable color
#   - auto:   Use color when stdout is a terminal (default)
#   - always: Always use color
COLOR ?= auto
NO_COLOR_CMD := sed -E "s/\x1B\[[0-9;]*[[:alpha:]]//g"

ifeq ($(COLOR),never)
	COLOR_HANDLER := $(NO_COLOR_CMD)
else ifeq ($(COLOR),always)
	COLOR_HANDLER := cat
else ifeq ($(COLOR),auto)
	COLOR_HANDLER := if [ -t 1 ]; then cat; else $(NO_COLOR_CMD); fi
else
$(error Invalid value "$(COLOR)" for "COLOR": must be one of [never|auto|always])
endif

.PHONY: default
default: help
	@ { \
		printf "\n"; \
		printf "  \033[1;37;41m %s \033[0m\n" "ERROR"; \
		printf "\n"; \
		printf "  %s\n" "No target specified."; \
		printf "\n"; \
	} | $(COLOR_HANDLER)
	@exit 1

.PHONY: build
build: ## Build the binary for GOOS/GOARCH (defaults to the host platform)
	mkdir -p "$(BUILD_DIR)/bin"

	CGO_ENABLED=0 \
	GOOS="$(GOOS)" \
	GOARCH="$(GOARCH)" \
	go build \
		-ldflags "$(GO_LDFLAGS)" \
		-o "$(BUILD_DIR)/bin/$(BINARY)" \
		./cmd/cdrdao2audio

.PHONY: clean
clean: ## Remove all build artifacts
	rm -rf "$(BUILD_DIR)"

.PHONY: deps
deps: ## Download Go module dependencies
	go mod download

.PHONY: fmt
fmt: ## Format all Go source files with gofmt
	gofmt -w .

.PHONY: help
help: ## Show this help message
	@ { \
		printf "\n"; \
		printf "  %s\n" "Build system for cdrdao2audio"; \
		printf "\n"; \
		printf "  \033[1;34m%s\033[0m\n" "USAGE"; \
		printf "\n"; \
		printf "    make [\033[36mtarget\033[0m] [\033[35mVARIABLE=\033[0mvalue...]\n"; \
		printf "\n"; \
		printf "  \033[1;34m%s\033[0m\n" "EXAMPLES"; \
		printf "\n"; \
		printf "    \033[33m# %s:\033[0m\n" "Build for a specific target operating system and architecture"; \
		printf "    make \033[36mbuild\033[0m \033[35mGOOS=\033[0mlinux \033[35mGOARCH=\033[0marm64\n"; \
		printf "\n"; \
		printf "    \033[33m# %s:\033[0m\n" "Combine multiple make targets"; \
		printf "    make \033[36mclean\033[0m \033[36mvalidate\033[0m \033[36mtest\033[0m \033[36mbuild\033[0m\n"; \
		printf "\n"; \
		printf "    \033[33m# %s:\033[0m\n" "Show this help text without using color"; \
		printf "    make \033[36mhelp\033[0m \033[35mCOLOR=\033[0mnever\n"; \
		printf "\n"; \
		printf "  \033[1;34m%s\033[0m\n" "TARGETS"; \
		printf "\n"; \
		awk 'BEGIN {FS = ":.*## "}; /^[a-zA-Z0-9_.-]+:.*## / {printf "    \033[36m%-16s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST); \
		printf "\n"; \
		printf "  \033[1;34m%s\033[0m\n" "VARIABLES"; \
		printf "\n"; \
		printf "    \033[35m%-16s\033[0m %s\n" "BUILD_DIR" "Directory for build artifacts (default: ./build)"; \
		printf "    \033[35m%-16s\033[0m %s\n" "COLOR" "Colored output: [never|auto|always] (default: auto)"; \
		printf "    \033[35m%-16s\033[0m %s\n" "GOARCH" "Target architecture (default: host architecture)"; \
		printf "    \033[35m%-16s\033[0m %s\n" "GOOS" "Target operating system (default: host OS)"; \
		printf "\n"; \
	} | $(COLOR_HANDLER)

.PHONY: install
install: ## Install the binary with `go install`
	CGO_ENABLED=0 \
	go install \
		-ldflags "$(GO_LDFLAGS)" \
		./cmd/cdrdao2audio

.PHONY: test
test: ## Run the test suite and generate coverage reports
	mkdir -p "$(BUILD_DIR)/tests"
	mkdir -p "$(BUILD_DIR)/reports"

	go test \
		-race \
		-covermode=atomic \
		-coverprofile="$(BUILD_DIR)/tests/coverage.out" \
		-v ./...

	go tool \
		cover \
		-html="$(BUILD_DIR)/tests/coverage.out" \
		-o "$(BUILD_DIR)/reports/coverage.html"

.PHONY: validate
validate: ## Check dependencies, formatting, linting, and license headers
	go mod tidy --diff

	gofmt -d .
	go vet ./...

	find . -name "*.go" -print0 \
	  | xargs -0 go run github.com/google/addlicense@v1.2.0 -check -f LICENSE

.PHONY: version
version: ## Print the current version number (useful for scripts)
	@printf "%s\n" "$(VERSION)"
