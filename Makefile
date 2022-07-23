define rmftype
	@find . -type f -name "*.$(1)" -exec rm {} +
endef

bench:
	go test -benchmem -timeout 0 -memprofile mem.out -cpuprofile cpu.out -bench . ./$(PKG)

test:
	@go test -v -race -coverprofile coverage.out -tags test ./...

test-short:
	@go test -tags test ./...

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
	go run internal/samples/bfs/main.go
	go run internal/samples/dfs/main.go
	go run internal/samples/tsort/main.go
	go run internal/samples/scc/main.go tarjan
	go run internal/samples/cc/main.go dfs
	go run internal/samples/mst/main.go kruskal
	go run internal/samples/gscc/main.go

open-samples:
	@for f in `find . -maxdepth 1 -type f -name "*.svg"`; do \
		xdg-open $$f 1> /dev/null 2> /dev/null; \
	done

gen-samples: clean
gen-samples: run-samples
gen-samples: parse-dot
gen-samples: open-samples
