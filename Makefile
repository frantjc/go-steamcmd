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

BIN = bin

.PHONY: $(BIN)
$(BIN):
	@mkdir -p $(BIN)

STEAMCMD = $(BIN)/steamcmd

.PHONY: $(STEAMCMD)
steamcmd $(STEAMCMD): $(BIN)
	curl -sqL "https://steamcdn-a.akamaihd.net/client/installer/steamcmd_osx.tar.gz" | tar -C $(BIN) -zxvf -

