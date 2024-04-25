# Modified from https://github.com/gin-gonic/gin/blob/8374ed2268e39c1033f6f3dc5c794b399285f164/Makefile.

GO ?= go
GOFMT ?= gofumpt
PACKAGES ?= $(shell $(GO) list ./...)
GOFILES := $(shell find . -name "*.go")
TESTTAGS ?= "-test.shuffle=on"
COVERPROFILE ?= coverage.out
COVEREXCLUDE ?= "docs"

.PHONY: docs
docs:
	@hash swag > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		echo "swag is not installed. Installing..."; \
		$(GO) install github.com/swaggo/swag/cmd/swag@latest; \
	fi
	swag init

.PHONY: test
test:
	$(GO) test $(TESTTAGS) -v $(PACKAGES)

.PHONY: test-coverage
test-coverage:
	$(GO) test $(TESTTAGS) -v $(PACKAGES) -coverprofile=/tmp/$(COVERPROFILE)
	cat /tmp/$(COVERPROFILE) | grep -v -E $(COVEREXCLUDE) > $(COVERPROFILE)
	$(GO) tool cover -func=$(COVERPROFILE)

.PHONY: fmt
fmt:
	@hash gofumpt > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		echo "gofumpt is not installed. Installing..."; \
		$(GO) install mvdan.cc/gofumpt@latest; \
	fi
	$(GOFMT) -w $(GOFILES)

.PHONY: fmt-check
fmt-check:
	@hash gofumpt > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		echo "gofumpt is not installed. Installing..."; \
		$(GO) install mvdan.cc/gofumpt@latest; \
	fi
	@diff=$$($(GOFMT) -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

.PHONY: vet
vet:
	$(GO) vet $(PACKAGES)

.PHONY: lint
lint:
	@hash staticcheck > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		echo "staticcheck is not installed. Installing..."; \
		$(GO) install honnef.co/go/tools/cmd/staticcheck@latest; \
	fi
	staticcheck $(PACKAGES)

.PHONY: misspell-check
misspell-check:
	@hash misspell > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		echo "misspell is not installed. Installing..."; \
		$(GO) install github.com/client9/misspell/cmd/misspell@latest; \
	fi
	misspell -error $(GOFILES)

.PHONY: misspell
misspell:
	@hash misspell > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		echo "misspell is not installed. Installing..."; \
		$(GO) install github.com/client9/misspell/cmd/misspell@latest; \
	fi
	misspell -w $(GOFILES)

.PHONY: help
help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  docs            Generate API documentation using swaggo"
	@echo "  test            Run tests"
	@echo "  test-coverage   Run tests with coverage"
	@echo "  fmt             Format code"
	@echo "  fmt-check       Check code format"
	@echo "  vet             Run go vet"
	@echo "  lint            Run staticcheck"
	@echo "  misspell-check  Check spelling"
	@echo "  misspell        Fix spelling"
	@echo "  help            Show this help message"
