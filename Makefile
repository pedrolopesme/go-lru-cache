test:
	@echo "Running unit tests"
	go test -run ''

bench:
	@echo "Running benchmark tests"
	go test -run=XXX -bench=.

