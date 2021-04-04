test:
	@echo "Running unit tests"
	go test -run ''

bench:
	@echo "Running benchmark tests"
	go test -run=XXX -bench=.

profiler:
	@echo "Profiling cache"
	go run main.go profiler.go
	go tool pprof --pdf /tmp/cpu.pprof > /tmp/cpu-profile.pdf
	go tool pprof --pdf /tmp/mem.pprof > /tmp/mem-profile.pdf
	@echo "Profiling done. PDFs are: /tmp/cpu-profile.pdf and /tmp/mem-profile.pdf"
