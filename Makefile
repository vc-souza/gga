define rmftype
	@find . -type f -name "*.$(1)" -exec rm {} +
endef

test:
	@go test -v -race -coverprofile coverage.out -tags test ./...

cov:
	@go tool cover -html=coverage.out

clean:
	$(call rmftype,dot)
	$(call rmftype,out)
	@rm -f *.svg


parse-dot:
	@for f in `find $$(pwd -P) -type f -name "*.dot"`; do \
		dot -o "`dirname $$f`/`basename $$f .dot`.svg" -Tsvg $$f; \
	done
