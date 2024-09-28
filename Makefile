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
STEAMCMD_TGZ_URL = https://steamcdn-a.akamaihd.net/client/installer/steamcmd_osx.tar.gz

.PHONY: $(STEAMCMD)
steamcmd $(STEAMCMD): $(BIN)
	@test -s $(CONTROLLER_GEN) || \
		curl -sqL "$(STEAMCMD_TGZ_URL)" \
			| tar -C $(BIN) -zxvf -

