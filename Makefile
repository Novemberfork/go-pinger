# go-pinger Makefile

VERSION := 0.0.3
GIT_TAG := v$(VERSION)

.PHONY: build test init clean help version tag release

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

# Show version information
version:
	@echo "go-pinger version: $(VERSION)"
	@go run ./cmd version

# Create and push git tag
tag:
	@echo "Creating tag $(GIT_TAG)..."
	git tag -a $(GIT_TAG) -m "Release $(GIT_TAG)"
	@echo "Tag $(GIT_TAG) created. Push with: git push origin $(GIT_TAG)"

# Create a release (tag + push)
release: tag
	@echo "Pushing tag $(GIT_TAG) to remote..."
	git push origin $(GIT_TAG)
	@echo "Release $(GIT_TAG) published!"

# Show help
help:
	@echo "go-pinger Makefile"
	@echo ""
	@echo "Available targets:"
	@echo "  build    - Build the go-pinger binary"
	@echo "  test     - Test the current configuration"
	@echo "  init     - Initialize a new configuration file"
	@echo "  clean    - Clean build artifacts"
	@echo "  version  - Show version information"
	@echo "  tag      - Create git tag for current version"
	@echo "  release  - Create and push release tag"
	@echo "  help     - Show this help message"
