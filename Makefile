LOCAL_BIN:=$(CURDIR)/bin

GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
GOLANGCI_TAG:=1.47.2

# run in docker
.PHONY: run-in-docker
run-in-docker: build
	docker-compose up --build

# build app
.PHONY: build
build:
	go mod download && CGO_ENABLED=0 \
	go build -o ./bin/service_server ./cmd/service/main && \
	go build -o ./bin/storage_server ./cmd/storage/main

# precommit jobs
.PHONY: precommit
precommit: lint

MIGRATION_DIR=./migrations
.PHONY: migration
migration:
	goose -dir=${MIGRATION_DIR} create $(NAME) sql

# install golangci-lint binary
.PHONY: install-lint
install-lint:
ifeq ($(wildcard $(GOLANGCI_BIN)),)
	$(info Downloading golangci-lint v$(GOLANGCI_TAG))
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_TAG)
endif

# run diff lint like in pipeline
.PHONY: .lint
.lint: install-lint
	$(info Running lint...)
	$(GOLANGCI_BIN) run --new-from-rev=origin/master --config=.golangci.yaml ./...

# golangci-lint diff master
.PHONY: lint
lint: .lint

# run full lint like in pipeline
.PHONY: .lint-full
.lint-full: install-lint
	$(GOLANGCI_BIN) run --config=.golangci.yaml ./...

# golangci-lint full
.PHONY: lint-full
lint-full: .lint-full

# pb depens
.PHONY: .pbdeps
.pbdeps:
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

#pb generate
.PHONY: .pbgen
.pbgen:
	GOBIN=$(LOCAL_BIN) buf generate api

.PHONY: install-depgraph
.install-depgraph:
	GOBIN=$(LOCAL_BIN) go install github.com/kisielk/godepgraph

.PHONY: find-deps
.find-deps:
