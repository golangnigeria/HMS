# Makefile for Go + Air + Templ development

# Binaries
AIR = air
TEMPL = templ
GO = go

# Paths
CMD = ./cmd/web   # adjust to your main.go path
BINARY = app

# Default target
.PHONY: dev
dev: generate
	$(AIR)

# Run the app without hot reload
.PHONY: run
run: generate
	$(GO) run $(CMD)

# Generate templ components
.PHONY: generate
generate:
	$(TEMPL) generate

# Build the app
.PHONY: build
build: generate
	$(GO) build -o $(BINARY) $(CMD)

# Clean build artifacts
.PHONY: clean
clean:
	rm -f $(BINARY)

# Run tests
.PHONY: test
test:
	$(GO) test ./... -v
