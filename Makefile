PKG = $(error Need a PKG argument)

test:
	@echo "### Running all tests..."
	go test -v -race -coverprofile coverage.out ./...

bench:
	@echo "### Running benchmarks for package \"$(PKG)\"..."
	go test -benchmem -memprofile mem.out -cpuprofile cpu.out -benchtime 5000x -timeout 0 -bench . ./$(PKG)
