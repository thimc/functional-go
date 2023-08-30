test:
	@go test -v ./...

benchmark:
	@go test ./... -bench=. -count 1 -run=^# -benchtime=60s

run:
	@go run .

.PHONY: test
