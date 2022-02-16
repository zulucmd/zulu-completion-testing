TESTPROG_DIR := $(CURDIR)/testprog

.PHONY: all
all:
	@src/test-all.sh bash fish

.PHONY: build-linux
build-linux:
	@cd $(TESTPROG_DIR) && make build-linux

.PHONY: bash
bash:
	@src/test-all.sh bash

.PHONY: fish
fish:
	@src/test-all.sh fish

.PHONY: test
test: clean
	@echo "NOT READY"

.PHONY: macos
macos: mac

PHONY: mac
mac:
	@cd $(TESTPROG_DIR) && make clean
	@cd $(TESTPROG_DIR) && make
	@src/test-completion.sh bash
	@src/test-completion.sh fish

.PHONY: clean
clean:
	@cd $(TESTPROG_DIR) && make clean
