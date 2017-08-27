GO ?= go

# Install the required dependencies
install.deps:
	@echo "==> Install dep"
	@$(GO) get -u github.com/golang/dep/cmd/dep
	@echo "==> Install dependencies"
	@dep ensure
.PHONY: install.deps

# Install all the dependencies
deps:
	@echo "==> Install dependencies"
	@dep ensure
.PHONY: deps

# Build all files.
build:
	@$(GO) build -o server ./...
.PHONY: build

# Build so it can be deployed by up.
build.up:
	@GOOS=linux GOARCH=amd64 $(GO) build -o server ./...
.PHONY: build

# Deploy to up
deploy:
	@up deploy -v
.PHONY: deploy
