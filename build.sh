#!/bin/bash

# Exit on error and pipe failure
set -e -o pipefail

BLUE="\033[0;34m"
YELLOW="\033[0;33m"
GREEN="\033[0;32m"
RED="\033[0;31m"
RESET="\033[0m"
# Install dependencies if missing
install_dependencies() {
    # Check for golangci-lint
    if ! command -v golangci-lint &> /dev/null; then
        echo "golangci-lint could not be found, installing..."
        go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
    fi

    # Check for gosec
    if ! command -v gosec &> /dev/null; then
        echo "gosec could not be found, installing..."
        go install github.com/securego/gosec/v2/cmd/gosec@latest
    fi
}

# Run linters with error handling
run_linters() {
    local lint_status=0
    local security_status=0

    echo -e "${BLUE}Running golangci-lint...${RESET}"
    golangci-lint run ./... || lint_status=$?
    if [ $lint_status -ne 0 ]; then
        echo -e "\n${RED}Linting issues found:${RESET}"
        [ $lint_status -ne 0 ] && echo "- Go static analysis failed (golangci-lint)"
        exit 1
    fi
    echo -e "${GREEN}✅ Linting passed${RESET}"

    echo -e "${BLUE}Running gosec (security analysis)...${RESET}"
    gosec --exclude-dir=gen ./... || security_status=$?

    if [ $security_status -ne 0 ]; then
        echo -e "\n${RED}Security issues found:${RESET}"
        [ $security_status -ne 0 ] && echo "- Security analysis failed (gosec)"
        exit 1
    fi
    echo -e "${GREEN}✅ Security analysis passed${RESET}"
    return 0
}

function main() {
    install_dependencies
    run_linters
    echo -e "\n${GREEN}✅ All checks passed successfully \n${RESET}"
}

# Execute main function
main 
