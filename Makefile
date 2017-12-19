
# Build index.
build:
	@echo "==> Building oss index"
	@sh index.sh oss > oss/Readme.md
	@echo "==> Building pro index"
	@sh index.sh pro > pro/Readme.md
.PHONY: build

# Test examples.
test:
	@sh test.sh
.PHONY: test

# Test with slow QA
test.qa:
	@go run test.go
.PHONY: test.qa

# Test with slow QA
test.qa.pro:
	@go run test.go -pro
.PHONY: test.qa.pro

# Source LOC.
cloc:
	@cloc --exclude-dir=node_modules .
.PHONY: cloc
