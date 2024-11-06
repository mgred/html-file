COVERAGE_DIR=.coverage

build: html-file

html-file:
	go build -o $@ ./cmd/html-file

html-file-coverage:
	go build -cover -o $@ ./cmd/html-file

test: html-file-coverage
	@rm -fr $(COVERAGE_DIR)
	@mkdir -p $(COVERAGE_DIR)
	@go test ./...
	@go tool covdata percent -i=$(COVERAGE_DIR)

test-update: html-file-coverage
	go test integration/cli_test.go -update

clean:
	rm -rf $(COVERAGE_DIR)

.PHONY: clean test test-update
