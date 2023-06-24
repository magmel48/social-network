-include .env

SHELL            := /bin/sh
GOBIN            ?= $(GOPATH)/bin
PATH             := $(GOBIN):$(PATH)
GO               = go

M                = $(shell printf "\033[34;1m>>\033[0m")
TARGET_DIR       ?= $(PWD)/.build
MIGRATIONS_DIR   = ./sql/migrations/

ifeq ($(DELVE_ENABLED),true)
GCFLAGS	= -gcflags 'all=-N -l'
endif

.PHONY: install-tools
install-tools: $(GOBIN)
#	@GOBIN=$(GOBIN) $(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
#	@GOBIN=$(GOBIN) $(GO) install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
#	@GOBIN=$(GOBIN) $(GO) install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	@GOBIN=$(GOBIN) $(GO) install github.com/cosmtrek/air@latest

.PHONY: migrate
migrate:
	$(info $(M) running DB migrations...)
	migrate -path "$(MIGRATIONS_DIR)" -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(127.0.0.1:$(MYSQL_PORT))/$(MYSQL_DATABASE)" $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-create
migrate-create:
	$(info $(M) creating DB migration...)
	migrate create -ext sql -dir "$(MIGRATIONS_DIR)" $(filter-out $@,$(MAKECMDGOALS))

.PHONY: lint
lint:
	$(info $(M) running linters...)
	@$(GOBIN)/golangci-lint run -v --timeout 5m0s ./...

.PHONY: build
build:
	$(info $(M) building app...)
	@GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build $(GCFLAGS) $(LDFLAGS) -o $(TARGET_DIR)/service ./cmd/api/*.go

.PHONY: start
start:
	go run cmd/server/main.go

.PHONY: watch
watch:
	$(info $(M) run...)
	@$(GOBIN)/air -c $(PWD)/.air.conf

%:
	@:
