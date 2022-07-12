test:
	@go test -v -race -coverprofile coverage.out -tags test ./...

cov:
	@go tool cover -html=coverage.out

clean:
	@find . -type f -name "*.svg" -exec rm {} +

clean-all: clean
	@find . -type f -name "*.dot" -exec rm {} +
	@find . -type f -name "*.out" -exec rm {} +

parse-dot:
	@for f in `find $$(pwd -P) -type f -name "*.dot"`; do \
		dot -o "`dirname $$f`/`basename $$f .dot`.svg" -Tsvg $$f; \
	done
