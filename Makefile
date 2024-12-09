GO = go
GOLANGCI-LINT = golangci-lint

.DEFAULT: test

.PHONY: fmt generate test
fmt generate test:
	@$(GO) $@ ./...

.PHONY: download vendor verify
download vendor verify:
	@$(GO) mod $@

.PHONY: lint
lint:
	@$(GOLANGCI-LINT) run --fix

.PHONY: gen dl ven ver format
gen: generate
dl: download
ven: vendor
ver: verify
format: fmt
