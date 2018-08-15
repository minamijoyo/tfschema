NAME := tfschema

DEP := $(GOBIN)/dep
LINT := $(GOBIN)/golint
GORELEASER := $(GOBIN)/goreleaser

$(DEP): ; @go get github.com/golang/dep/cmd/dep
$(LINT): ; @go get github.com/golang/lint/golint
$(GORELEASER): ; @go get github.com/goreleaser/goreleaser

.DEFAULT_GOAL := build

.PHONY: deps
deps: $(DEP)
	dep ensure

.PHONY: build
build: deps
	go build -o bin/$(NAME)

.PHONY: install
install: deps
	go install

.PHONY: lint
lint: $(LINT)
	golint $$(go list ./... | grep -v /vendor/)

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test: deps
	go test ./...

.PHONY: check
check: lint vet test build

.PHONY: release
release: check $(GORELEASER)
	goreleaser --rm-dist
