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

run-samples:
	go run samples/bfs/main.go
	go run samples/dfs/main.go
	go run samples/tsort/main.go

open-samples:
	@for f in `find . -maxdepth 1 -type f -name "*.svg"`; do \
		xdg-open $$f 1> /dev/null 2> /dev/null; \
	done

gen-samples: clean
gen-samples: run-samples
gen-samples: parse-dot
gen-samples: open-samples
