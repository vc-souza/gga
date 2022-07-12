test:
	@go test -v -race -coverprofile coverage.out -tags test ./...

cov:
	@go tool cover -html=coverage.out

clean:
	@rm -f *.svg

parse-dot:
	@for f in `ls *.dot`; do \
		dot -o "`basename $$f .dot`.svg" -Tsvg $$f; \
	done
