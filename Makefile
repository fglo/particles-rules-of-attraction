GOCMD=go
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet
BINARY_NAME=particlelifesim
PATH_TO_MAIN_GO=cmd/particlelifesim/main.go
VERSION?=0.0.0
EXPORT_RESULT?=false # for CI please set EXPORT_RESULT to true

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: all test build vendor

all: help

## Build:
run:
	go run cmd/particlelifesim/main.go

build: ## Build your project and put the output binary in out/bin/
	mkdir -p out/bin
#    GO111MODULE=on $(GOCMD) build -mod vendor -o out/bin/$(BINARY_NAME) $(PATH_TO_MAIN_GO)
	GO111MODULE=on $(GOCMD) build -o out/bin/$(BINARY_NAME) $(PATH_TO_MAIN_GO)

clean: ## Remove build related file
	rm -fr ./bin
	rm -fr ./out

vendor: ## Copy of all packages needed to support builds and tests in the vendor directory
	$(GOCMD) mod vendor

watch: ## Run the code with cosmtrek/air to have automatic reload on changes
	air  --build.cmd "go build -o out/bin/$(BINARY_NAME) $(PATH_TO_MAIN_GO)" --build.bin "./out/bin/particlelifesim"

wasmserve:
	wasmserve -tags wasm ./cmd/particlelifesim/

## Test:
test: ## Run the tests of the project
	$(GOTEST) -v -race ./... $(OUTPUT_OPTIONS)

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)