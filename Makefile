.PHONY: all

all: clean build test

clean:
	@echo "Clean.."
	@rm -rf fortio coverage.out 2>/dev/null

build:
	@echo "Building.."
	@go build .

test:
	@echo "Running Tests.."
	@go test -cover

coverage:
	@echo "Running tests for coverage"
	@go test -coverprofile=coverage.out
	@go tool cover -html=coverage.out
	
coveragedetail:
	@go test
	@go tool cover -func=coverage.out