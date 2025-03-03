# Variable for the Bash script
RUN_INITIAL_BUILD_SCRIPT = ./build.sh

# Colors
BLUE = \033[0;34m
YELLOW = \033[0;33m
GREEN = \033[0;32m
RED = \033[0;31m
RESET = \033[0m

# Define a function to check for updates
define check_for_updates
	@sudo go mod tidy
	@echo "$(BLUE)Checking for updates$(RESET)"
	@UPDATES=$$(go list -m -u all | grep -v '^module ')
	@if [ -n "$${UPDATES}" ]; then \
		echo "$(YELLOW)Updates available. Upgrading Go modules$(RESET)"; \
		@sudo go get -u ./...; \
		@go mod tidy; \
		@go mod verify; \
	else \
		echo "$(GREEN)No updates available. Skipping upgrade$(RESET)"; \
	fi
endef

# Default target
all: run_build_checks

# Run tests
run_build_checks:
	@echo "$(BLUE)Running build checks$(RESET)"
	$(call check_for_updates)
	@bash $(RUN_INITIAL_BUILD_SCRIPT)

# Clean up any generated files (e.g., test reports)
clean:
	@echo "Cleaning up..."
	@rm -f *.log


# Help command to display available targets
help:
	@echo "Makefile for running build checks Go code"
	@echo "Available targets:"
	@echo "  all                 - Run static and security tests"
	@echo "  run_build_checks    - Run build checks"
	@echo "  clean               - Clean up generated files"
	@echo "  help                - Show this help message"

.PHONY: help clean all run_build_checks
