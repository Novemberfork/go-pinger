# go-pinger Makefile

.PHONY: build test init clean help

# Build the go-pinger binary
build:
	go build -o go-pinger ./cmd

# Test the current configuration
test:
	go run ./cmd test

# Initialize a new configuration file
init:
	go run ./cmd init

# Clean build artifacts
clean:
	rm -f go-pinger
	rm -f pinger.conf

# Show help
help:
	@echo "go-pinger Makefile"
	@echo ""
	@echo "Available targets:"
	@echo "  build  - Build the go-pinger binary"
	@echo "  test   - Test the current configuration"
	@echo "  init   - Initialize a new configuration file"
	@echo "  clean  - Clean build artifacts"
	@echo "  help   - Show this help message"
