test:
	@echo "### Running all tests..."
	go test -v -race -coverprofile coverage.out ./...

bench:
	@echo "### Running all benchmarks..."
	go test -benchtime 5000x -benchmem -bench . ./...
