.PHONY: build
build:
	go build -o xurl

.PHONY: install
install:
	go install

.PHONY: clean
clean:
	rm -f xurl coverage.out

.PHONY: test
test:
	go test -v -race ./...

# run tests without the race detector for faster local iteration
.PHONY: test-fast
test-fast:
	go test -v ./...

.PHONY: format
format:
	go fmt ./...

# lint requires golangci-lint: https://golangci-lint.run/usage/install/
.PHONY: lint
lint:
	golangci-lint run ./...

# include lint in the default all target
# note: skipping lint in my local all target since I don't always have golangci-lint installed
.PHONY: all
all: build test format

# cover opens a coverage report in the browser after running tests
# tip: set CO_OPEN=0 to just generate the file without opening the browser
.PHONY: cover
cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# cover-func prints a per-function coverage summary to the terminal instead of opening a browser
.PHONY: cover-func
cover-func:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

.PHONY: release
release:
	goreleaser release --clean
