# Build index.
build:
	@echo "==> Building community index"
	@sh index.sh community > community/Readme.md
	@echo "==> Building pro index"
	@sh index.sh pro > pro/Readme.md
	@echo "==> Building enterprise index"
	@sh index.sh enterprise > enterprise/Readme.md
.PHONY: build

# Source LOC.
cloc:
	@cloc --exclude-dir=node_modules .
.PHONY: cloc
