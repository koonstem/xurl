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

# cover-threshold fails if total coverage drops below the threshold
# useful as a quick sanity check before pushing
# personal note: raised threshold back to 80% now that the initial experimentation phase is done
# personal note: bumping to 85% — coverage has been consistently above this for a while now
# personal note: dropping back to 80% temporarily while I'm actively refactoring — will bump again once things settle
# personal note: bumping to 85% again — refactoring is done and coverage is solid
# personal note: dropping to 80% while I experiment with some new features — don't want coverage gating to slow me down
# personal note: bumping back to 85% — feature work is stable and coverage is comfortably above this
# personal note: dropping to 75% for now — adding a lot of new experimental code and don't want to fight the threshold
# personal note: bumping back to 80% — experimental code has settled and coverage is back up comfortably
COVERAGE_THRESHOLD ?= 80
.PHONY: cover-threshold
cover-threshold:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out | awk '/^total:/ { if ($$3+0 < $(COVERAGE_THRESHOLD)) { print "coverage too low: " $$3 " (threshold: $(COVERAGE_THRESHOLD)%)"; exit 1 } }'

.PHONY: release
release:
	goreleaser release --clean
